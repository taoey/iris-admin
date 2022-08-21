package api

import (
	"github.com/kataras/iris/v12"
)

type Controller interface {
	Bind(ctx iris.Context) error
	Check(ctx iris.Context) error
	Run(ctx iris.Context) (interface{}, error)
}

// if req.Method == http.MethodGet {
// 	// TODO 反射自动绑定参数
// 	c.Req = ctx.URLParams()
// 	return nil
// }
// if err := ctx.ReadJSON(&c.Req); err != nil {
// 	return err
// }
