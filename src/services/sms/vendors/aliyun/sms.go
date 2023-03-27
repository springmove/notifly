package aliyun

import (
	"encoding/json"
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
)

type SMS struct {
	base.BaseSMSProvider
	clients map[string]*dysmsapi.Client
}

func (s *SMS) Init() {
	s.clients = map[string]*dysmsapi.Client{}
	for name, endpoint := range s.Endpoints {
		client, err := dysmsapi.NewClientWithAccessKey(endpoint.Region, endpoint.AppKey, endpoint.AppSecret)
		if err != nil {
			sptty.Log(sptty.ErrorLevel, err.Error(), base.Aliyun)
			continue
		}

		s.clients[name] = client
	}
}

func (s *SMS) Send(req *base.ReqSMS) error {
	ep, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	client, exist := s.clients[req.Endpoint]
	if !exist {
		return errors.New("Client Not Found ")
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = req.Mobile
	request.SignName = ep.SignName
	request.TemplateCode = ep.TemplateCode

	content, _ := json.Marshal(req.Content)
	request.TemplateParam = string(content)

	resp, err := client.SendSms(request)
	if err != nil {
		return err
	}

	if resp.Code != "OK" {
		errBoby, _ := json.Marshal(resp)
		return errors.New(string(errBoby))
	}

	return nil
}
