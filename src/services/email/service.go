package email

import (
	"errors"

	"github.com/linshenqi/sptty"
	"gopkg.in/gomail.v2"
)

const (
	ServiceName = "email"
)

type Request struct {
	Endpoint string   `json:"endpoint"`
	MailTo   []string `json:"mail_to"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
}

type Service struct {
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

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}

func (s *Service) ServiceName() string {
	return ServiceName
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

func (s *Service) Send(req Request) error {

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
