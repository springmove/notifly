package goeasy

import (
	"fmt"
	"github.com/springmove/notifly/src/base"
	"testing"
)

func getIM() *IM {
	im := IM{}
	if err := im.Init(&base.IMEndpoint{
		Provider: base.IMGoEasy,
		AppKey:   "BC-0304a351925c472ea61f3c16b32c77af",
		Hosts:    []string{"https://rest-hangzhou.goeasy.io/publish"},
	}); err != nil {
		return nil
	}

	return &im
}

func TestIM(t *testing.T) {
	im := getIM()
	if im == nil {
		return
	}

	host := im.GetHostByRegion("")
	if err := im.PostMessage(&base.IMMessage{
		Channel: "1",
		Host:    host,
		Content: "awefawef",
	}); err != nil {
		fmt.Print(err.Error())
	}

}
