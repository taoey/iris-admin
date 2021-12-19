package test

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/taoey/iris-admin/pkg/sys/req"
	"github.com/taoey/iris-admin/pkg/sys/resp"
)

func MapParmsHandler(ctx iris.Context) {
	params := req.Read2Map(ctx)
	fmt.Println(params)
	age := params["tao"]
	fmt.Println(age)
	ctx.JSON(resp.OkResponse())
}

func ErrorHandler(ctx iris.Context) {
	a := ctx.URLParamDefault("a", "10")
	b := ctx.URLParamDefault("b", "2")
	int_a, _ := strconv.Atoi(a)
	int_b, _ := strconv.Atoi(b)
	c := int_a / int_b
	ctx.JSON(resp.OkResponseWithRet(c))
}
