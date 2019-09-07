package notify

import (
	"github.com/kataras/iris"
	"github.com/linshenqi/sptty"
)

type Controllers struct {
	service *Service
}

// 发送模板消息
func (s *Controllers) postTemplateMsg(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := TemplateMsg{}

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_REQ, err.Error()))
		return
	}

	err = s.service.wechat.SendTemplateMsg(req.Endpoint, req.Touser, req.TemplateID, req.MiniProgram.Page, req.FormID, req.Data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_MSG, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
}

// 发送自定义客服消息
func (s *Controllers) postCustomerMsg(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := CustomerMsg{}

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_REQ, err.Error()))
		return
	}

	err = s.service.wechat.SendCustomerMsg(req.Endpoint, req.Body)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_MSG, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
}

func (s *Controllers) postEnterpriseMsg(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := EnterpriseMsg{}

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_REQ, err.Error()))
		return
	}

	err = s.service.wechat.SendEnterpriseGroupMsg(req.Endpoint, req.ChatID, req.MsgType, req.Safe, req.Text.Content)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(NOTIFY_ERR_MSG, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
}
