package email

import (
	"fmt"
	"testing"

	"github.com/springmove/notifly/src/base"
)

var endpoint = "ashibro"

func getService() *Service {
	email := Service{cfg: Config{Endpoints: map[string]Endpoint{
		endpoint: {
			Author: "notify@ashibro.com",
			Pwd:    "css199520.",
			Host:   "smtp.mxhichina.com",
			Port:   25,
		},
	}}}

	email.load()
	return &email
}

func TestService(t *testing.T) {
	email := getService()
	if err := email.Send(&base.ReqEmail{
		Endpoint: endpoint,
		MailTo:   []string{"linshenqi@outlook.com"},
		Subject:  "test",
		Body:     "ashibro test",
	}); err != nil {
		fmt.Println(err.Error())
	}
}
