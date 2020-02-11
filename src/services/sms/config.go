package sms

type Config struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Provider     string `yaml:"provider"`
	AppKey       string `yaml:"app_key"`
	AppSecret    string `yaml:"app_secret"`
	Region       string `yaml:"region"`
	SignName     string `yaml:"sign_name"`
	TemplateCode string `yaml:"template_code"`
	HostNum      string `yaml:"host_num"`
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
