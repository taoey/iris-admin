package api

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

type Controller interface {
	Bind(ctx iris.Context) error
	Check(ctx iris.Context) error
	Run(ctx iris.Context) (interface{}, error)
}

type DefaultController struct {
	Req interface{}
}

func (c *DefaultController) Bind(ctx iris.Context) error {
	req := ctx.Request()
	if req.Method == http.MethodGet {
		// TODO 反射自动绑定参数
		c.Req = ctx.URLParams()
		return nil
	}
	if err := ctx.ReadJSON(&c.Req); err != nil {
		return err
	}
	return nil
}

func (c *DefaultController) Check(ctx iris.Context) error {
	return nil
}

func (c *DefaultController) Run(ctx iris.Context) (interface{}, error) {

	return nil, nil
}
