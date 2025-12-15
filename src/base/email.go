package base

import (
	"gopkg.in/gomail.v2"
)

const (
	ServiceEmail = "email"
)

var Email IServiceEmail

type IServiceEmail interface {
	EmailClient(index ...int) *EmailClient
}

type ReqEmail struct {
	MailTo  []string
	Subject string
	Body    string
}

type EmailEntry struct {
	Author string `yaml:"Author"`
	Sender string `yaml:"Sender"`
	Pwd    string `yaml:"Pwd"`
	Host   string `yaml:"Host"`
	Port   int    `yaml:"Port"`
}

type EmailClient struct {
	Dialer   *gomail.Dialer
	Endpoint *EmailEntry
}

func (s *EmailClient) Send(req *ReqEmail) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.Endpoint.Author+"<"+s.Endpoint.Sender+">")
	msg.SetHeader("To", req.MailTo...)
	msg.SetHeader("Subject", req.Subject)
	msg.SetBody("text/html", req.Body)

	return s.Dialer.DialAndSend(msg)
}
