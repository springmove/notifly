package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type Service struct {
	app sptty.Sptty
	cfg Config

	http *resty.Client
}

func (s *Service) Init(service sptty.Sptty) error {
	s.app = service

	_ = s.app.GetConfig("wechat", &s.cfg)
	s.http = sptty.CreateHttpClient(&sptty.HttpClientConfig{
		Timeout:      8,
		PushInterval: 1,
		MaxRetry:     3,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return nil
}

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}

func (s *Service) getAccessToken(endpoint string) (string, error) {
	ep := s.cfg.Endpoints[endpoint]
	tr := TokenResp{}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		ep.AppID,
		ep.AppSecret)

	resp, err := s.http.R().Get(url)

	if err != nil {
		return "", err
	} else {
		err = json.Unmarshal(resp.Body(), &tr)
		if err != nil {
			return "", err
		} else {
			if tr.Errcode == 0 {
				return tr.AccessToken, nil
			} else {
				return "", errors.New(fmt.Sprintf("errcode:%d errmsg:%s", tr.Errcode, tr.Errmsg))
			}
		}
	}
}

func (s *Service) SendTemplateMsg(endpoint string, openid string, templateid string, page string, formid string, data interface{}) error {
	token, err := s.getAccessToken(endpoint)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)

	r := s.http.R().SetBody(&MsgTemplate{
		Touser:     openid,
		TemplateID: templateid,
		//Page:            page,
		MiniProgram: MiniProgram{
			AppID: s.cfg.Endpoints["mp"].AppID,
			Page:  page,
		},
		//FormID:          formid,
		Data: data,
		//EmphasisKeyword: "",
	}).SetHeader("content-type", "application/json")

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

func (s *Service) SendCustomerMsg(endpoint string, openid string, content string) error {
	token, err := s.getAccessToken(endpoint)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", token)
	r := s.http.R().SetBody(&MsgCustomer{
		Touser:  openid,
		Msgtype: "text",
		Text:    content,
	}).SetHeader("content-type", "application/json")

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
