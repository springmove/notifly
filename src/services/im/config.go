package im

import "github.com/linshenqi/notifly/src/base"

type Config struct {
	Endpoints map[string]base.IMEndpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return ServiceName
}

func (s *Config) Validate() error {
	return nil
}

func (s *Config) Default() interface{} {
	return &Config{}
}
