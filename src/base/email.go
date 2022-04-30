package base

const (
	ServiceEmail = "email"
)

type ReqEmail struct {
	Endpoint string   `json:"endpoint"`
	MailTo   []string `json:"mail_to"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
}

type IServiceEmail interface {
	Send(req *ReqEmail) error
}
