package wechat

type Config struct {
	Enable    bool                            `yaml:"enable"`
	ResPath   string                          `yaml:"respath"`
	Endpoints map[string]ConfigWecharEndpoint `yaml:"endpoints"`
}

type ConfigWecharEndpoint struct {
	AppID     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
}
