package wechat

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

type MsgTemplate struct {
	Touser          string      `json:"touser"`
	TemplateID      string      `json:"template_id"`
	Page            string      `json:"page"`
	FormID          string      `json:"form_id"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}

type TemplateValue struct {
	Value string `json:"value"`
}
