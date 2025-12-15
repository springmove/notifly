package email

import (
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
	"gopkg.in/gomail.v2"
)

type Service struct {
	sptty.BaseService

	cfg     Config
	clients []*base.EmailClient
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	if !s.cfg.Enable {
		sptty.Log(sptty.InfoLevel, "Service Disabled", s.ServiceName())
		return nil
	}

	if err := s.initClients(); err != nil {
		return nil
	}

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceEmail
}

func (s *Service) initClients() error {
	s.clients = []*base.EmailClient{}
	for k, v := range s.cfg.Configs {
		dialer := gomail.NewDialer(v.Host, v.Port, v.Author, v.Pwd)
		s.clients = append(s.clients, &base.EmailClient{
			Dialer:   dialer,
			Endpoint: &s.cfg.Configs[k],
		})
	}

	return nil
}

func (s *Service) EmailClient(index ...int) *base.EmailClient {
	target := 0
	if len(index) > 0 {
		target = index[0]
	}

	return s.clients[target]
}
