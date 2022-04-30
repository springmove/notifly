package rongcloud

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/linshenqi/notifly/src/base"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type Resp struct {
	Code      int    `json:"code"`
	SessionID string `json:"sessionId"`
}

type SMS struct {
	base.BaseSMSProvider
	http *resty.Client
}

func (s *SMS) Init() {
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *SMS) Send(req *base.ReqSMS) error {
	ep, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	nonce := fmt.Sprintf("%d", rand.Intn(9999))
	ts := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	sign := generateSign(ep.AppSecret, nonce, ts)

	r := s.http.R().
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetHeader("App-Key", ep.AppKey).
		SetHeader("Nonce", nonce).
		SetHeader("Timestamp", ts).
		SetHeader("Signature", sign)

	vals := content2ValuesFormat(req.Content)
	url := "http://api.sms.ronghub.com/sendNotify.json"
	body := fmt.Sprintf("mobile=%s&templateId=%s&region=%s&%s", req.Mobile, ep.TemplateCode, ep.Region, vals)

	resp, err := r.SetBody(body).Post(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%+v", resp)
	}

	respBody := Resp{}
	if err := json.Unmarshal(resp.Body(), &respBody); err != nil {
		return err
	}

	if respBody.Code != http.StatusOK {
		return fmt.Errorf("%+v", respBody)
	}

	return nil
}

func content2ValuesFormat(content map[string]string) string {
	vals := []string{}
	index := 0
	for k := range content {
		index++
		vals = append(vals, fmt.Sprintf("p%d=%s", index, strings.TrimSpace(content[k])))
	}

	return strings.Join(vals, "&")
}

func generateSign(secret string, nonce string, ts string) string {
	sign := sptty.Sha1(fmt.Sprintf("%s%s%s", secret, nonce, ts))
	return sign
}
