package notify

import (
	"github.com/springmove/notifly/src/base"
	"github.com/springmove/notifly/src/services/wechat"
	"github.com/springmove/sptty"
)

const (
	ServiceName = "notify"
)

type Service struct {
	wechat *wechat.Service
}

func (s *Service) Init(app sptty.ISptty) error {
	s.wechat = app.GetService(base.ServiceWechatNotify).(*wechat.Service)

	app.AddRoute("POST", "/customer-image", s.postCustomerImage)
	app.AddRoute("POST", "/customer-msgs", s.postCustomerMsg)
	app.AddRoute("POST", "/template-msgs", s.postTemplateMsg)
	app.AddRoute("POST", "/mp-template-msgs", s.postMPTemplateMsg)
	app.AddRoute("POST", "/enterprise-msgs", s.postEnterpriseMsg)
	app.AddRoute("POST", "/mp-subscribe-msgs", s.postMPSubMsg)

	return nil
}

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}

func (s *Service) ServiceName() string {
	return ServiceName
}
