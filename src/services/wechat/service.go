package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kataras/iris/core/errors"
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/sptty"
	"gopkg.in/resty.v1"
)

type Service struct {
	sptty.BaseService

	http *resty.Client
}

func (s *Service) Init(app sptty.ISptty) error {

	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceWechatNotify
}

// func (s *Service) getEnterpriseAccessToken(endpoint string) (string, error) {
// 	ep := s.cfg.Endpoints[endpoint]
// 	tr := TokenResp{}

// 	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
// 		ep.AppID,
// 		ep.AppSecret)

// 	resp, err := s.http.R().Get(url)

// 	if err != nil {
// 		return "", err
// 	} else {
// 		err = json.Unmarshal(resp.Body(), &tr)
// 		if err != nil {
// 			return "", err
// 		} else {
// 			if tr.Errcode == 0 {
// 				return tr.AccessToken, nil
// 			} else {
// 				return "", errors.New(fmt.Sprintf("errcode:%d errmsg:%s", tr.Errcode, tr.Errmsg))
// 			}
// 		}
// 	}
// }

// func (s *Service) SendEnterpriseGroupMsg(endpoint string, chatID string, msgType string, safe int, content string) error {
// 	token, err := s.getEnterpriseAccessToken(endpoint)
// 	if err != nil {
// 		return err
// 	}

// 	url := fmt.Sprintf(" https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s", token)

// 	body := EnterpriseGroupMsg{
// 		ChatID:  chatID,
// 		MsgType: msgType,
// 		Safe:    safe,
// 		Text: Content{
// 			Content: content,
// 		},
// 	}

// 	r := s.http.R().SetBody(&body).SetHeader("content-type", "application/json")

// 	resp, err := r.Post(url)

// 	msgresp := MsgResp{}

// 	if err != nil {
// 		return err
// 	} else {
// 		if resp.StatusCode() != http.StatusOK {
// 			return errors.New(fmt.Sprintf("%d", resp.StatusCode()))
// 		} else {
// 			err := json.Unmarshal(resp.Body(), &msgresp)
// 			if err != nil {
// 				return err
// 			}

// 			if msgresp.Errcode == 0 {
// 				return nil
// 			} else {
// 				return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
// 			}
// 		}
// 	}
// }

// func (s *Service) SendTemplateMsg(token string, openid string, templateid string, page string, formid string, data interface{}) error {

// 	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)

// 	body := MsgTemplate{
// 		Touser:     openid,
// 		TemplateID: templateid,
// 		//Page:            page,
// 		MiniProgram: MiniProgram{
// 			AppID: s.cfg.Endpoints["mp"].AppID,
// 			Page:  page,
// 		},
// 		//FormID:          formid,
// 		Data: data,
// 		//EmphasisKeyword: "",
// 	}
// 	r := s.http.R().SetBody(&body).SetHeader("content-type", "application/json")

// 	resp, err := r.Post(url)

// 	msgresp := MsgResp{}

// 	if err != nil {
// 		return err
// 	} else {
// 		if resp.StatusCode() != http.StatusOK {
// 			return errors.New(fmt.Sprintf("%d", resp.StatusCode()))
// 		} else {
// 			err := json.Unmarshal(resp.Body(), &msgresp)
// 			if err != nil {
// 				return err
// 			}

// 			if msgresp.Errcode == 0 {
// 				return nil
// 			} else {
// 				return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
// 			}
// 		}
// 	}
// }

func (s *Service) SendMPTemplateMsg(token string, openid string, templateid string, page string, formid string, data interface{}) error {

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", token)

	body := MPMsgTemplate{
		Touser:     openid,
		TemplateID: templateid,
		Data:       data,
		FormID:     formid,
	}
	r := s.http.R().SetBody(&body).SetHeader("content-type", "application/json")

	resp, err := r.Post(url)

	msgresp := MsgResp{}

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(fmt.Sprintf("%d", resp.StatusCode()))
		} else {
			err := json.Unmarshal(resp.Body(), &msgresp)
			if err != nil {
				return err
			}

			if msgresp.Errcode == 0 {
				return nil
			} else {
				return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
			}
		}
	}
}

//func (s *Service) SendCustomerMsg(endpoint string, openid string, content string) error {
//	token, err := s.getAccessToken(endpoint)
//	if err != nil {
//		return err
//	}
//
//	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", token)
//	r := s.http.R().SetBody(&MsgCustomer{
//		Touser:  openid,
//		Msgtype: "text",
//		Text:    content,
//	}).SetHeader("content-type", "application/json")
//
//	resp, err := r.Post(url)
//
//	msgresp := MsgResp{}
//
//	if err != nil {
//		return err
//	} else {
//		if resp.StatusCode() != http.StatusOK {
//			return errors.New(fmt.Sprintf("%d", resp.StatusCode()))
//		} else {
//			err := json.Unmarshal(resp.Body(), &msgresp)
//			if err != nil {
//				return err
//			}
//
//			if msgresp.Errcode == 0 {
//				return nil
//			} else {
//				return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
//			}
//		}
//	}
//}

//func (s *Service) EnterpriseRobot

// func (s *Service) UploadImage(token string, file string) (string, error) {

// 	fullPath := path.Join(s.cfg.ResPath, file)
// 	image, err := ioutil.ReadFile(fullPath)
// 	if err != nil {
// 		return "", err
// 	}

// 	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, "image")

// 	r := s.http.R().SetFileReader("image", file, bytes.NewReader(image))

// 	resp, err := r.Post(url)

// 	msgresp := MsgRespImage{}

// 	if err != nil {
// 		return "", err
// 	} else {
// 		if resp.StatusCode() != http.StatusOK {
// 			return "", errors.New(fmt.Sprintf("%d", resp.StatusCode()))
// 		} else {
// 			err := json.Unmarshal(resp.Body(), &msgresp)
// 			if err != nil {
// 				return "", err
// 			}

// 			if msgresp.Errcode == 0 {
// 				return msgresp.MediaID, nil
// 			} else {
// 				return "", errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
// 			}
// 		}
// 	}
// }

func (s *Service) SendCustomerMsg(token string, body interface{}) error {

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", token)
	r := s.http.R().SetBody(body).SetHeader("content-type", "application/json")

	resp, err := r.Post(url)

	msgresp := MsgResp{}

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(fmt.Sprintf("%d", resp.StatusCode()))
		} else {
			err := json.Unmarshal(resp.Body(), &msgresp)
			if err != nil {
				return err
			}

			if msgresp.Errcode == 0 {
				return nil
			} else {
				return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", msgresp.Errcode, msgresp.Errmsg))
			}
		}
	}
}
