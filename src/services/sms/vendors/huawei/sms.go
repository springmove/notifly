package huawei

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/linshenqi/notifly/src/base"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type SMS struct {
	base.BaseSMSProvider
	http *resty.Client
}

func (s *SMS) Init() {
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *SMS) Send(req base.Request) error {
	ep, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	appKey := ep.AppKey
	appSecret := ep.AppSecret
	from := ep.HostNum
	to := req.Mobile
	templateId := ep.TemplateCode
	templateParas := req.Content["code"]

	//用户信息加密
	passwordDigest, nonce, created := GetWSSE(appSecret)
	xWSSEParam := fmt.Sprintf(`UsernameToken Username="%v",PasswordDigest="%v",Nonce="%v",Created="%v"`, appKey, passwordDigest, nonce, created)

	//设置头部信息
	r := s.http.R().
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`).
		SetHeader("X-WSSE", xWSSEParam)

	//发送请求
	url := "https://rtcsms.cn-north-1.myhuaweicloud.com:10743/sms/batchSendSms/v1"
	body := fmt.Sprintf(`from=%v&to=%v&templateId=%v&templateParas="[\"%v\"]"`, from, to, templateId, templateParas)

	resp, err := r.SetBody(body).Post(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%+v", resp)
	}

	return nil
}

func GetWSSE(appSecret string) (string, int, string) {
	Created := time.Now().Format("2006-01-02T15:04:05Z") //创建一个时间戳，并转成W3DTF格式
	Nonce := rand.Intn(1000000000)                       //产生一个随机数

	//sha256sum 就是linux上计算sha256值的一个程序。通过sha256sum 可以算出目标的sha256值。
	sha := sha256.Sum256([]byte(fmt.Sprintf("%v%v%v", Nonce, Created, appSecret)))
	encodingStr := base64.StdEncoding.EncodeToString(sha[:])

	return string(encodingStr), Nonce, Created
}
