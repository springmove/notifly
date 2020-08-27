package wechat

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/linshenqi/sptty"
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

func getSrv() *Service {
	return &Service{
		http: sptty.CreateHttpClient(sptty.DefaultHttpClientConfig()),
	}
}

func TestBotMsg(t *testing.T) {
	srv := getSrv()
	if err := srv.PostGroupBotMsg(&ReqBotMsg{
		BotKey: "87b201af-0a79-4904-827b-f0b85a7e286d",
		BotMsgContent: BotMsgContent{
			MsgType: BotMsgTypeText,
			Text: &Content{
				Content: "test text",
			},
		},
	}); err != nil {
		fmt.Println(err.Error())
	}

	if err := srv.PostGroupBotMsg(&ReqBotMsg{
		BotKey: "87b201af-0a79-4904-827b-f0b85a7e286d",
		BotMsgContent: BotMsgContent{
			MsgType: BotMsgTypeMarkdown,
			Markdown: &Content{
				Content: "*test markdown*",
			},
		},
	}); err != nil {
		fmt.Println(err.Error())
	}
}

func TestMiniProgramSubscribeMsg(t *testing.T) {
	srv := getSrv()

	accessToken, err := srv.GetAccessToken("wx77d83a2aa6c324ab", "2c6ecb6fe8a0394715704149a6afc56b")
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
