package im

import (
	"github.com/linshenqi/notifly/src/base"
	"github.com/linshenqi/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]base.IMEndpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return base.ServiceIM
}
