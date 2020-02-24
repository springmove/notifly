package sms

import (
	"fmt"
	"github.com/linshenqi/notifly/src/services/base"
	"testing"
)

var endpoint = "ashibro"
var endpointCN = "ashibro_cn"

func getService() *Service {
	sms := Service{cfg: Config{Endpoints: map[string]base.Endpoint{
		endpoint: {
			Provider:  base.Twilio,
			AppKey:    "ACc1921dda2551f493a630645923b24d17",
			AppSecret: "b218b891ef85c68b48ee53d84b9c8b42",
			HostNum:   "+12055761160",
		},
		endpointCN: {
			Provider:     base.Aliyun,
			AppKey:       "LTAI4FqKBiAkx65UoYqx8fiS",
			AppSecret:    "ZVJbGx3TDQPYHMNmlHtLjQySisu5dX",
			Region:       "cn-hangzhou",
			SignName:     "Ashibro",
			TemplateCode: "SMS_183261852",
		},
	}}}

	sms.setupProviders()
	sms.initProviders()

	return &sms
}

func TestService(t *testing.T) {
	sms := getService()
	//if err := sms.Send(base.Request{
	//	Endpoint: endpoint,
	//	Mobile:   "+8618049956365",
	//	Content:  "ashibro test sms2",
	//}); err != nil {
	//	fmt.Println(err.Error())
	//}

	if err := sms.Send(base.Request{
		Endpoint: endpointCN,
		Mobile:   "18621182783",
		Content:  fmt.Sprintf("{\"code\":\"1234\"}"),
	}); err != nil {
		fmt.Println(err.Error())
	}
}
