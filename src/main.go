package main

import (
	"flag"
	"github.com/linshenqi/notifly/src/services/notify"
	"github.com/linshenqi/notifly/src/services/wechat"
	"github.com/linshenqi/sptty"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.ConfFromFile(*cfg)

	services := map[string]sptty.Service{
		"wechat": &wechat.Service{},
		"notify": &notify.Service{},
		//"auth":   &auth.AuthService{},
		//"jwt":    &jwt.JwtService{},
	}

	configs := sptty.SpttyConfig{
		"http":   sptty.HttpConfig{},
		"model":  sptty.ModelConfig{},
		"wechat": wechat.Config{},
		//"jwt":    jwt.JwtConfig{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
