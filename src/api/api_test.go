package api

import (
	"fmt"
	"testing"

	"github.com/springmove/notifly/src/base"
	"github.com/stretchr/testify/assert"
)

func TestCustomerMsg(t *testing.T) {
	cfg := Config{
		Url: "http://127.0.0.1:10002",
	}

	notifly := Notifly{}
	_ = notifly.InitService(&cfg)

	mediaID, err := notifly.PostCustomerImage(&base.CustomerImage{
		Endpoint: "mp",
		Path:     "/home/linshenqi/Pictures/wyy_qrcode.jpg",
	})

	assert.Nil(t, err)

	err = notifly.PostCustomerMsg(&base.ReqCustomerMsg{
		Endpoint: "mp",
		Body: base.CustomerMsgImage{
			CustomerMsg: base.CustomerMsg{
				ToUser:  "oHdMv5aqHTw56H56G4dfedPEGRVk",
				MsgType: "image",
			},
			Image: base.MsgImage{
				MediaID: mediaID,
			},
		},
	})

	assert.Nil(t, err)

	//err = notifly.PostCustomerMsg(&notify.CustomerMsg{
	//	Endpoint: "mp",
	//	Body: wechat.CustomerMsgText{
	//		CustomerMsg: wechat.CustomerMsg{
	//			ToUser:  "oHdMv5aqHTw56H56G4dfedPEGRVk",
	//			MsgType: "text",
	//		},
	//		Text: wechat.Content{
	//			Content: "终于等到你，长按扫码关注我们的公众号开启专属你的宠物消息推送吧！",
	//		},
	//	},
	//})

	//resp.Data.(wechat.WXAuthResponse)
	//fmt.Printf("%s\n", wx)
}

func TestTemplateMsg(t *testing.T) {
	cfg := Config{
		Url: "http://127.0.0.1:10002",
	}

	notifly := Notifly{}
	_ = notifly.InitService(&cfg)

	err := notifly.PostTemplateMsg(&base.TemplateMsg{
		Endpoint: "mp",
		MsgTemplate: base.MsgTemplate{
			Touser:     "",
			TemplateID: "",
			FormID:     "",
			Data: map[string]base.TemplateValue{
				"keyword1": {
					Value: "通过",
				},
				"keyword2": {
					Value: "0000-00-00",
				},
				"keyword3": {
					Value: "通过",
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)

	//resp.Data.(wechat.WXAuthResponse)
	//fmt.Printf("%s\n", wx)
}
