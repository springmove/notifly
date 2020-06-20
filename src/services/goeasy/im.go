package goeasy

import (
	"errors"
	"fmt"
	"github.com/linshenqi/notifly/src/services/base"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type IM struct {
	base.BaseIMProvider
	http *resty.Client
}

func (s *IM) Init(endpoint *base.IMEndpoint) error {
	if err := s.BaseIMProvider.Init(endpoint); err != nil {
		return err
	}

	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
	return nil
}

func (s *IM) GetHostByRegion(code string) string {
	defaultHost := s.BaseIMProvider.GetEndpoint().Hosts[0]

	// todo: select host by code
	return defaultHost
}

func (s *IM) PostMessage(msg *base.IMMessage) error {
	endpoint := s.BaseIMProvider.GetEndpoint()
	url := fmt.Sprintf("%s?appkey=%s&channel=%s&content=%s", msg.Host, endpoint.AppKey, msg.Channel, msg.Content)
	resp, err := s.http.R().Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("PostMessage Failed: %d", resp.StatusCode()))
	}

	return nil
}
