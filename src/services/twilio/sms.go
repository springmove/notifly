package twilio

import (
	"errors"
	"github.com/linshenqi/notifly/src/services/sms"
	"github.com/sfreiberg/gotwilio"
)

type SMS struct {
	sms.BaseSMSProvider
	clients map[string]*gotwilio.Twilio
}

func (s *SMS) Init() {
	s.clients = map[string]*gotwilio.Twilio{}
	for name, endpoint := range s.Endpoints {
		client := gotwilio.NewTwilioClient(endpoint.AppKey, endpoint.AppSecret)
		s.clients[name] = client
	}
}

func (s *SMS) Send(req sms.Request) error {
	ep, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	client, exist := s.clients[req.Endpoint]
	if !exist {
		return errors.New("Client Not Found ")
	}

	resp, _, err := client.SendSMS(ep.HostNum, req.Mobile, req.Content, "", "")
	if err != nil {
		return err
	}

	if resp.Status != "sent" {
		return errors.New(resp.Status)
	}

	return nil
}
