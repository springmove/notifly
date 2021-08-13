package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linshenqi/notifly/src/services/notify"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type Config struct {
	Url string `yaml:"url"`
}

type Notifly struct {
	cfg  *Config
	http *resty.Client
}

func (s *Notifly) InitService(cfg *Config) error {

	s.cfg = cfg
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())

	return nil
}

func (s *Notifly) PostCustomerMsg(req *notify.CustomerMsg) error {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/customer-msgs", s.cfg.Url)
	resp, err := r.Post(url)

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(string(resp.Body()))
		} else {
			return nil
		}
	}
}

func (s *Notifly) PostTemplateMsg(req *notify.TemplateMsg) error {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/template-msgs", s.cfg.Url)
	resp, err := r.Post(url)

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(string(resp.Body()))
		} else {
			return nil
		}
	}
}

func (s *Notifly) PostEnterpriseMsg(req *notify.EnterpriseMsg) error {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/enterprise-msgs", s.cfg.Url)
	resp, err := r.Post(url)

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(string(resp.Body()))
		} else {
			return nil
		}
	}
}

func (s *Notifly) PostCustomerImage(req *notify.CustomerImage) (string, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/customer-image", s.cfg.Url)
	resp, err := r.Post(url)

	image := notify.CustomerImageResp{}
	if err != nil {
		return "", err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return "", errors.New(string(resp.Body()))
		} else {
			_ = json.Unmarshal(resp.Body(), &image)
			return image.MediaID, nil
		}
	}
}

func (s *Notifly) PostMPSubscribeMsg(req *notify.MPSubscribeMsg) error {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/mp-subscribe-msgs", s.cfg.Url)
	resp, err := r.Post(url)

	if err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New(string(resp.Body()))
		} else {
			return nil
		}
	}
}
