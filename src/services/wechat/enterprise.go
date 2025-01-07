package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/springmove/notifly/src/base"
)

// 企业微信群机器人消息推送
func (s *Service) PostGroupBotMsg(req *base.ReqBotMsg) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", req.BotKey)
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
