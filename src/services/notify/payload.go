package notify

import "github.com/linshenqi/notifly/src/services/wechat"

const (
	NOTIFY_ERR_REQ = "0001"
	NOTIFY_ERR_MSG = "0002"
)

type CustomerMsg struct {
	Endpoint string `json:"endpoint"`
	OpenID   string `json:"openid"`
	Content  string `json:"content"`
}

type TemplateMsg struct {
	Endpoint string `json:"endpoint"`
	wechat.MsgTemplate
}
