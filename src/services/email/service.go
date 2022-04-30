package email

import (
	"errors"

	"github.com/linshenqi/notifly/src/base"
	"github.com/linshenqi/sptty"
	"gopkg.in/gomail.v2"
)

type Service struct {
	sptty.BaseService

	cfg     Config
	dialers map[string]*gomail.Dialer
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	s.load()
	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceEmail
}

func (s *Service) load() {
	s.dialers = map[string]*gomail.Dialer{}
	for name, endpoint := range s.cfg.Endpoints {
		dialer := gomail.NewDialer(endpoint.Host, endpoint.Port, endpoint.Author, endpoint.Pwd)
		s.dialers[name] = dialer
	}
}

func (s *Service) getEndpoint(endpoint string) (*Endpoint, error) {
	ep, exist := s.cfg.Endpoints[endpoint]
	if !exist {
		return nil, errors.New("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *Service) Send(req *base.ReqEmail) error {

	endpoint, err := s.getEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	dialer, exist := s.dialers[req.Endpoint]
	if !exist {
		return errors.New("Dialer Not Found ")
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", req.Endpoint+"<"+endpoint.Author+">")
	msg.SetHeader("To", req.MailTo...)
	msg.SetHeader("Subject", req.Subject)
	msg.SetBody("text/html", req.Body)

	return dialer.DialAndSend(msg)
}
