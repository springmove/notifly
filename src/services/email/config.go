package email

type Config struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Author string `yaml:"author"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
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
