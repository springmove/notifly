package im

import (
	"errors"
	"fmt"
	"github.com/linshenqi/notifly/src/services/base"
	"github.com/linshenqi/notifly/src/services/goeasy"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "im"
)

type Service struct {
	cfg       Config
	providers map[string]base.IIMProvider
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

func (s *Service) initProviders() {
	s.providers = map[string]base.IIMProvider{}

	var imProvider base.IIMProvider
	endpoints := s.cfg.Endpoints

	for k, v := range endpoints {
		switch v.Provider {
		case base.IMGoEasy:
			imProvider = &goeasy.IM{}
			imEndpoint := endpoints[k]
			if err := imProvider.Init(&imEndpoint); err != nil {
				sptty.Log(sptty.ErrorLevel, fmt.Sprintf("Init IMProvider Error: %s", err.Error()), s.ServiceName())
				return
			}

		default:
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("IMProvider Error: %s", v.Provider), s.ServiceName())
			return
		}

		s.providers[k] = imProvider
	}
}

func (s *Service) getProvider(providerName string) (base.IIMProvider, error) {
	provider, exist := s.providers[providerName]
	if !exist {
		return nil, errors.New("Provider Not Found ")
	}

	return provider, nil
}

func (s *Service) PostMessage(providerName string, msg *base.IMMessage) error {
	provider, err := s.getProvider(providerName)
	if err != nil {
		return err
	}

	return provider.PostMessage(msg)
}

func (s *Service) GetHostByRegion(providerName string, code string) (string, error) {
	provider, err := s.getProvider(providerName)
	if err != nil {
		return "", err
	}

	host := provider.GetHostByRegion(code)
	return host, nil
}
