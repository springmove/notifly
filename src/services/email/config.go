package email

import (
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Author string `yaml:"author"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

func (s *Config) ConfigName() string {
	return base.ServiceEmail
}
