package email

import (
	"fmt"
	"testing"

	"github.com/springmove/notifly/src/base"
)

var email base.IServiceEmail

func getService() base.IServiceEmail {

	if email == nil {
		srv := Service{cfg: Config{Configs: []base.EmailEntry{
			{
				Author: "kqdigital",
				Sender: "linshenqi@springmove.net",
				Pwd:    "",
				Host:   "smtp.exmail.qq.com",
				Port:   465,
			},
		}}}

		_ = srv.initClients()

		email = &srv
	}

	return email
}

func TestService(t *testing.T) {
	email := getService()
	if err := email.EmailClient().Send(&base.ReqEmail{
		MailTo:  []string{"linshenqi@outlook.com"},
		Subject: "kqdigital into test",
		Body:    "kqdigital into test",
	}); err != nil {
		fmt.Println(err.Error())
	}
}
