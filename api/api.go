package api

import (
	"errors"
	"fmt"
	"github.com/linshenqi/notifly/src/services/notify"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type NotiflyConfig struct {
	Url          string            `yaml:"url"`
	Timeout      int               `yaml:"timeout"`
	Headers      map[string]string `yaml:"headers"`
	PushInterval int               `yaml:"push_interval"`
	MaxRetry     int               `yaml:"max_retry"`
}

type Notifly struct {
	cfg  *NotiflyConfig
	http *resty.Client
}

func (s *Notifly) InitService(cfg *NotiflyConfig) error {

	s.cfg = cfg

	clientCfg := sptty.HttpClientConfig{
		Timeout:      s.cfg.Timeout,
		Headers:      s.cfg.Headers,
		PushInterval: s.cfg.PushInterval,
		MaxRetry:     s.cfg.MaxRetry,
	}

	s.http = sptty.CreateHttpClient(&clientCfg)

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
