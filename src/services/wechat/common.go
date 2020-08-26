package wechat

import (
	"encoding/json"
	"fmt"
)

func (s *Service) GetAccessToken(appID string, appSecret string) (string, error) {
	tr := TokenResp{}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		appID,
		appSecret)

	resp, err := s.http.R().Get(url)

	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(resp.Body(), &tr); err != nil {
		return "", err
	}

	if tr.Errcode != 0 {
		return "", fmt.Errorf("%+v", tr)
	}

	return tr.AccessToken, nil
}
