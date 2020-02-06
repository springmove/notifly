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

	services := sptty.Services{
		&wechat.Service{},
		&notify.Service{},
	}

	configs := sptty.Configs{
		&wechat.Config{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
