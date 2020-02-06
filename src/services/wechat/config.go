package wechat

type Config struct {
	ResPath   string              `yaml:"respath"`
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

func (s *Config) ConfigName() string {
	return ServiceName
}

func (s *Config) Validate() error {
	return nil
}

func (s *Config) Default() interface{} {
	return &Config{
		ResPath: "./",
	}
}
