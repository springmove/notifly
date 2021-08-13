package base

import "errors"

const (
	IMGoEasy = "goeasy"
)

type IMEndpoint struct {
	Provider string   `yaml:"provider"`
	AppKey   string   `yaml:"app_key"`
	Hosts    []string `yaml:"hosts"`
}

type IMMessage struct {
	Channel string
	Content string
	Host    string
}

type IIMProvider interface {
	GetHostByRegion(code string) string
	Init(endpoint *IMEndpoint) error
	PostMessage(msg *IMMessage) error
	GetEndpoint() *IMEndpoint
}

type BaseIMProvider struct {
	endpoint *IMEndpoint
	IIMProvider
}

func (s *BaseIMProvider) Init(endpoint *IMEndpoint) error {
	if len(endpoint.Hosts) == 0 {
		return errors.New("At Least One Host Should Be Set ")
	}

	s.endpoint = endpoint
	return nil
}

func (s *BaseIMProvider) GetEndpoint() *IMEndpoint {
	return s.endpoint
}

func (s *BaseIMProvider) GetHostByRegion(code string) string {
	return ""
}

func (s *BaseIMProvider) PostMessage(msg *IMMessage) error {
	return nil
}
