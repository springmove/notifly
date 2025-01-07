package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/springmove/notifly/src/base"
)

type MsgValue struct {
	Value interface{} `json:"value"`
}

type SubscribeMsg struct {
	Touser           string              `json:"touser"`
	TemplateID       string              `json:"template_id"`
	Page             string              `json:"page,omitempty"`
	MiniprogramState string              `json:"miniprogram_state,omitempty"`
	Lang             string              `json:"lang,omitempty"`
	Data             map[string]MsgValue `json:"data"`
}

type ReqMiniProgramSubscribeMsg struct {
	AccessToken string
	SubscribeMsg
}

// 小程序订阅消息
func (s *Service) SendMiniProgramSubscribeMsg(req *ReqMiniProgramSubscribeMsg) error {

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", req.AccessToken)

	body, _ := json.Marshal(req.SubscribeMsg)
	r := s.http.R().SetBody(body).SetHeader("content-type", "application/json")

	msgresp := base.MsgResp{}
	resp, err := r.Post(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%+v", resp)
	}

	if err := json.Unmarshal(resp.Body(), &msgresp); err != nil {
		return err
	}

	if msgresp.Errcode != 0 {
		return fmt.Errorf("%+v", msgresp)
	}

	return nil
}
