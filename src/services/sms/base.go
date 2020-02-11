package sms

import (
	"errors"
)

const (
	Aliyun = "aliyun"
	Twilio = "twilio"
)

type Request struct {
	Provider string `json:"provider"`
	Endpoint string `json:"endpoint"`
	Mobile   string `json:"mobile"`
	Content  string `json:"content"`
}

type ISMSProvider interface {
	Send(req Request) error
	Init()
	GetEndpoint(name string) (*Endpoint, error)
	AddEndpoint(name string, endpoint Endpoint)
}

type BaseSMSProvider struct {
	Endpoints map[string]Endpoint
}

func (s *BaseSMSProvider) GetEndpoint(name string) (*Endpoint, error) {
	ep, exist := s.Endpoints[name]
	if !exist {
		return nil, errors.New("Sms Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *BaseSMSProvider) Init() {
}

func (s *BaseSMSProvider) AddEndpoint(name string, endpoint Endpoint) {
	if s.Endpoints == nil {
		s.Endpoints = map[string]Endpoint{}
	}
	s.Endpoints[name] = endpoint
}
