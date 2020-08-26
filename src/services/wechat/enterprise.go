package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BotMsgTypeText     = "text"
	BotMsgTypeMarkdown = "markdown"
)

type BotMsgContent struct {
	MsgType  string   `json:"msgtype"`
	Text     *Content `json:"text,omitempty"`
	Markdown *Content `json:"markdown,omitempty"`
}

type ReqBotMsg struct {
	BotKey string
	BotMsgContent
}

// 企业微信群机器人消息推送
func (s *Service) PostGroupBotMsg(req *ReqBotMsg) error {
	url := fmt.Sprintf("http://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", req.BotKey)
	body, _ := json.Marshal(req.BotMsgContent)
	r := s.http.R().SetBody(body).SetHeader("content-type", "application/json")
	resp, err := r.Post(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%+v", resp)
	}

	return nil
}
