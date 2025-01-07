package base

const (
	ServiceWechatNotify = "wechat-notify"
)

const (
	BotMsgTypeText     = "text"
	BotMsgTypeMarkdown = "markdown"
)

const (
	ErrReqNotify = "ErrReqNotify"
	ErrMsgNotify = "ErrMsgNotify"
)

type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type MsgCustomer struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Text    string `json:"text"`
}

type MsgResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MsgRespImage struct {
	MsgResp
	MediaID string `json:"media_id"`
}

type MsgTemplate struct {
	Touser          string      `json:"touser"`
	TemplateID      string      `json:"template_id"`
	Url             string      `json:"url"`
	MiniProgram     MiniProgram `json:"miniprogram"`
	FormID          string      `json:"form_id"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}

type MPMsgTemplate struct {
	Touser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	//Url             string      `json:"url"`
	//MiniProgram     MiniProgram `json:"miniprogram"`
	FormID string      `json:"form_id"`
	Data   interface{} `json:"data"`
	//EmphasisKeyword string      `json:"emphasis_keyword"`
}

type TemplateValue struct {
	Value string `json:"value"`
}

type MiniProgram struct {
	AppID string `json:"appid"`
	Page  string `json:"page"`
}

type Content struct {
	Content string `json:"content"`
}

type EnterpriseGroupMsg struct {
	ChatID  string  `json:"chatid"`
	MsgType string  `json:"msgtype"`
	Safe    int     `json:"safe"`
	Text    Content `json:"text"`
}

type CustomerMsg struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

type MsgImage struct {
	MediaID string `json:"media_id"`
}

type MsgLink struct {
	Title    string `json:"title"`
	Desc     string `json:"description"`
	Url      string `json:"url"`
	ThumbUrl string `json:"thumb_url"`
}

type CustomerMsgImage struct {
	CustomerMsg
	Image MsgImage `json:"image"`
}

type CustomerMsgLink struct {
	CustomerMsg
	Link MsgLink `json:"link"`
}

type CustomerMsgText struct {
	CustomerMsg
	Text Content `json:"text"`
}

type BotMsgContent struct {
	MsgType  string   `json:"msgtype"`
	Text     *Content `json:"text,omitempty"`
	Markdown *Content `json:"markdown,omitempty"`
}

type ReqBotMsg struct {
	BotKey string
	BotMsgContent
}

type CustomerImage struct {
	Endpoint string `json:"endpoint"`
	Path     string `json:"path"`
}

type CustomerImageResp struct {
	MediaID string `json:"media_id"`
}

type ReqCustomerMsg struct {
	Endpoint string      `json:"endpoint"`
	Body     interface{} `json:"body"`
}

type TemplateMsg struct {
	Endpoint string `json:"endpoint"`
	MsgTemplate
}

type EnterpriseMsg struct {
	Endpoint string `json:"endpoint"`
	EnterpriseGroupMsg
}

type MPSubscribeMsg struct {
	Endpoint string `json:"endpoint"`
	// wechat.MsgSub
}

type IServiceWechatNotify interface {
	PostGroupBotMsg(req *ReqBotMsg) error
}
