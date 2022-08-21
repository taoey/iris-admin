package api

import (
	"github.com/kataras/iris/v12"
	"github.com/taoey/iris-admin/pkg/idl/pb_file/wx"
)

func NewWxCheckController(ctx iris.Context) Controller {
	return &WxCheckController{}
}

type WxCheckController struct {
	// Req *wx.WxCheckRequest
	Req map[string]string
}

func (c *WxCheckController) Bind(ctx iris.Context) error {
	c.Req = ctx.URLParams()
	return nil
}

func (c *WxCheckController) Check(ctx iris.Context) error {
	return nil
}

func (c *WxCheckController) Run(ctx iris.Context) (interface{}, error) {
	resp := wx.WxCheckResponse{
		Echostr: c.Req["echostr"],
	}
	return &resp, nil
}
