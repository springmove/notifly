package email

import (
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Configs []base.EmailEntry `yaml:"Configs"`
}

func (s *Config) ConfigName() string {
	return base.ServiceEmail
}
