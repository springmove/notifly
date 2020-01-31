package notify

import "github.com/linshenqi/notifly/src/services/wechat"

const (
	NOTIFY_ERR_REQ = "0001"
	NOTIFY_ERR_MSG = "0002"
)

type CustomerImage struct {
	Endpoint string `json:"endpoint"`
	Path     string `json:"path"`
}

type CustomerImageResp struct {
	MediaID string `json:"media_id"`
}

type CustomerMsg struct {
	Endpoint string      `json:"endpoint"`
	Body     interface{} `json:"body"`
}

type TemplateMsg struct {
	Endpoint string `json:"endpoint"`
	wechat.MsgTemplate
}

type EnterpriseMsg struct {
	Endpoint string `json:"endpoint"`
	wechat.EnterpriseGroupMsg
}

type MPSubscribeMsg struct {
	Endpoint string `json:"endpoint"`
	wechat.MsgSub
}
