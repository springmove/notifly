package api

import (
	"fmt"
	"github.com/linshenqi/notifly/src/services/notify"
	"github.com/linshenqi/notifly/src/services/wechat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerMsg(t *testing.T) {
	cfg := NotiflyConfig{
		Url:          "http://127.0.0.1:10002",
		Timeout:      3,
		Headers:      map[string]string{"Content-Type": "application/json"},
		PushInterval: 1,
		MaxRetry:     1,
	}

	notifly := Notifly{}
	notifly.InitService(&cfg)

	mediaID, err := notifly.PostCustomerImage(&notify.CustomerImage{
		Endpoint: "mp",
		Path:     "/home/linshenqi/Pictures/wyy_qrcode.jpg",
	})

	assert.Nil(t, err)

	err = notifly.PostCustomerMsg(&notify.CustomerMsg{
		Endpoint: "mp",
		Body: wechat.CustomerMsgImage{
			CustomerMsg: wechat.CustomerMsg{
				ToUser:  "oHdMv5aqHTw56H56G4dfedPEGRVk",
				MsgType: "image",
			},
			Image: wechat.MsgImage{
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
	cfg := NotiflyConfig{
		Url:          "http://127.0.0.1:10002",
		Timeout:      3,
		Headers:      map[string]string{"Content-Type": "application/json"},
		PushInterval: 1,
		MaxRetry:     1,
	}

	notifly := Notifly{}
	notifly.InitService(&cfg)

	err := notifly.PostTemplateMsg(&notify.TemplateMsg{
		Endpoint: "mp",
		MsgTemplate: wechat.MsgTemplate{
			Touser:     "oHdMv5aqHTw56H56G4dfedPEGRVk",
			TemplateID: "W5OwhKlQCWy7Cf9s9MgxO-Ytsk86MQ_KJYinYBQLsT4",
			FormID:     "",
			Data: map[string]wechat.TemplateValue{
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
