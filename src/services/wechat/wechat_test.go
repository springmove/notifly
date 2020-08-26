package wechat

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWechat(t *testing.T) {
	body, _ := json.Marshal(ReqBotMsg{
		BotKey: "wef",
		BotMsgContent: BotMsgContent{
			MsgType: BotMsgTypeText,
			Markdown: &Content{
				Content: "awef",
			},
		},
	})
	fmt.Println(string(body))

	body, _ = json.Marshal(ReqMiniProgramSubscribeMsg{
		AccessToken: "t1",
		SubscribeMsg: SubscribeMsg{
			Touser:     "uid",
			TemplateID: "t1",
			Page:       "wef/wef/wef/wef",
			Data: map[string]MsgValue{
				"kwef2": {
					Value: "wef",
				},
				"23": {
					Value: "3334",
				},
			},
		},
	})
	fmt.Println(string(body))
}
