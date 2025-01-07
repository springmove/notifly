package wechat

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
)

func TestWechat(t *testing.T) {
	body, _ := json.Marshal(base.ReqBotMsg{
		BotKey: "wef",
		BotMsgContent: base.BotMsgContent{
			MsgType: base.BotMsgTypeText,
			Markdown: &base.Content{
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

func getSrv() *Service {
	return &Service{
		http: sptty.CreateHttpClient(sptty.DefaultHttpClientConfig()),
	}
}

func TestBotMsg(t *testing.T) {
	srv := getSrv()
	if err := srv.PostGroupBotMsg(&base.ReqBotMsg{
		BotKey: "0a9ee1c9-7db0-47ac-9c0b-ab0b69b6155b",
		BotMsgContent: base.BotMsgContent{
			MsgType: base.BotMsgTypeText,
			Text: &base.Content{
				Content: "test text",
			},
		},
	}); err != nil {
		fmt.Println(err.Error())
	}

	if err := srv.PostGroupBotMsg(&base.ReqBotMsg{
		BotKey: "0a9ee1c9-7db0-47ac-9c0b-ab0b69b6155b",
		BotMsgContent: base.BotMsgContent{
			MsgType: base.BotMsgTypeMarkdown,
			Markdown: &base.Content{
				Content: "*test markdown*",
			},
		},
	}); err != nil {
		fmt.Println(err.Error())
	}
}

func TestMiniProgramSubscribeMsg(t *testing.T) {
	srv := getSrv()

	accessToken, err := srv.GetAccessToken("", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := srv.SendMiniProgramSubscribeMsg(&ReqMiniProgramSubscribeMsg{
		AccessToken: accessToken,
		SubscribeMsg: SubscribeMsg{
			Touser:     "o3Q5b5Je_zyzyGUazgaMC6dMe9Pw",
			TemplateID: "w7jnRnZJRnxJJqKQFhU3uLlfawmYwRQkuWRDT_T00fE",
			Data: map[string]MsgValue{
				"phrase1": {
					Value: "测试",
				},
				"thing2": {
					Value: "状态以改变，awefwaef",
				},
				"thing3": {
					Value: "测试备注提示",
				},
			},
		},
	}); err != nil {
		fmt.Println(err.Error())
		return
	}
}
