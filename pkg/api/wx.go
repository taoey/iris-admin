package api

import (
	"github.com/kataras/iris/v12"
	"github.com/taoey/iris-admin/pkg/idl/pb_file/user"
)

func NewWxCheckController(ctx iris.Context) Controller {
	return &WxCheckController{
		DefaultController{&user.LoginReq{}},
	}
}

type WxCheckController struct {
	DefaultController
}

func (c *WxCheckController) Check(ctx iris.Context) error {
	return nil
}

func (c *WxCheckController) Run(ctx iris.Context) (interface{}, error) {
	return c.Req, nil
}
