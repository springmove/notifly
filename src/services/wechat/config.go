package wechat

type Config struct {
	Enable    bool                            `yaml:"enable"`
	Endpoints map[string]ConfigWecharEndpoint `yaml:"endpoints"`
}

type ConfigWecharEndpoint struct {
	AppID     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
}
