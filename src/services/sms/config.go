package sms

import (
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]base.Endpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return base.ServiceSMS
}
