package sms

import (
	"errors"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "sms"
)

type Service struct {
	cfg       Config
	providers map[string]ISMSProvider
}

func (s *Service) Init(app sptty.Sptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	s.initProviders()

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

func (s *Service) Send(req Request) error {
	provider, err := s.getProvider(req.Provider)
	if err != nil {
		return err
	}

	return provider.Send(req)
}

func (s *Service) initProviders() {
	for k, v := range s.cfg.Endpoints {
		provider, err := s.getProvider(v.Provider)
		if err != nil {
			continue
		}

		provider.AddEndpoint(k, v)
	}

	for _, provider := range s.providers {
		provider.Init()
	}
}

func (s *Service) getProvider(providerType string) (ISMSProvider, error) {
	provider, exist := s.providers[providerType]
	if !exist {
		return nil, errors.New("Provider Not Found ")
	}

	return provider, nil
}

func (s *Service) SetupProviders(providers map[string]ISMSProvider) {
	s.providers = providers
}
