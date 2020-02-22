package twilio

import (
	"errors"
	"github.com/linshenqi/notifly/src/services/base"
	"github.com/sfreiberg/gotwilio"
)

type SMS struct {
	base.BaseSMSProvider
	clients map[string]*gotwilio.Twilio
}

func (s *SMS) Init() {
	s.clients = map[string]*gotwilio.Twilio{}
	for name, endpoint := range s.Endpoints {
		client := gotwilio.NewTwilioClient(endpoint.AppKey, endpoint.AppSecret)
		s.clients[name] = client
	}
}

func (s *SMS) Send(req base.Request) error {
	ep, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	client, exist := s.clients[req.Endpoint]
	if !exist {
		return errors.New("Client Not Found ")
	}

	resp, ex, err := client.SendSMS(ep.HostNum, req.Mobile, req.Content, "", "")
	if err != nil {
		return err
	}

	if ex != nil {
		return ex
	}

	if resp.Status != "sent" {
		return errors.New(resp.Status)
	}

	return nil
}
