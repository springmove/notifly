package sms

import (
	"fmt"
	"testing"

	"github.com/springmove/notifly/src/base"
)

// var endpoint = "ashibro"
// var endpointCN = "ashibro_cn"
var endpointHuawei = "huawei"

func getService() *Service {
	sms := Service{cfg: Config{Endpoints: map[string]base.Endpoint{
		//endpoint: {
		//	Provider:  base.Twilio,
		//	AppKey:    "ACc1921dda2551f493a630645923b24d17",
		//	AppSecret: "b218b891ef85c68b48ee53d84b9c8b42",
		//	HostNum:   "+12055761160",
		//},
		//endpointCN: {
		//	Provider:     base.RongCloud,
		//	AppKey:       "",
		//	AppSecret:    "",
		//	Region:       "86",
		//	SignName:     "Ashibro",
		//	TemplateCode: "bT0pus3s4CD8w4LrgZ_g2e",
		//},
		endpointHuawei: {
			Provider:     base.Huawei,
			AppKey:       "",
			AppSecret:    "",
			TemplateCode: "",
			HostNum:      "",
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
	//	Content: map[string]string{
	//				"code": "1234",
	//			},
	//}); err != nil {
	//	fmt.Println(err.Error())
	//}

	if err := sms.Send(&base.ReqSMS{
		Endpoint: endpointHuawei,
		Mobile:   "18621182783",
		Content: map[string]string{
			"code": "012345",
		},
	}); err != nil {
		fmt.Println(err.Error())
	}
}
