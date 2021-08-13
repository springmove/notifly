package sms

import (
	"errors"
	"github.com/linshenqi/notifly/src/base"
	"github.com/linshenqi/notifly/src/services/sms/vendors/aliyun"
	"github.com/linshenqi/notifly/src/services/sms/vendors/huawei"
	"github.com/linshenqi/notifly/src/services/sms/vendors/rongcloud"
	"github.com/linshenqi/notifly/src/services/sms/vendors/twilio"
	"github.com/linshenqi/sptty"
)

const ()

type Service struct {
	cfg       Config
	providers map[string]base.ISMSProvider
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	s.setupProviders()
	s.initProviders()

	return nil
}

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}

func (s *Service) ServiceName() string {
	return base.ServiceSMS
}

func (s *Service) Send(req base.Request) error {
	endpoint, err := s.getEndpoint(req.Endpoint)
	if err != nil {
		return err
	}

	provider, err := s.getProvider(endpoint.Provider)
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

func (s *Service) getProvider(providerType string) (base.ISMSProvider, error) {
	provider, exist := s.providers[providerType]
	if !exist {
		return nil, errors.New("Provider Not Found ")
	}

	return provider, nil
}

func (s *Service) getEndpoint(endpoint string) (*base.Endpoint, error) {
	ep, exist := s.cfg.Endpoints[endpoint]
	if !exist {
		return nil, errors.New("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *Service) setupProviders() {
	s.providers = map[string]base.ISMSProvider{
		base.Aliyun:    &aliyun.SMS{},
		base.Twilio:    &twilio.SMS{},
		base.RongCloud: &rongcloud.SMS{},
		base.Huawei:    &huawei.SMS{},
	}
}
