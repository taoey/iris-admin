package api

import (
	"github.com/kataras/iris/v12"
)

type Controller interface {
	Bind(ctx iris.Context) error
	Check(ctx iris.Context) error
	Run(ctx iris.Context) (interface{}, error)
}
