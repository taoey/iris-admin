package router

import (
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/taoey/iris-admin/pkg/api"
)

//	Message 类常量
const (
	MESSAGE_OK    = 200
	MESSAGE_ERROR = 500
)

/**
 * @Description: 接收ctx中的json参数，转化为map
 * @param ctx
 * @return map[string]interface{}
 */
func Read2Map(ctx iris.Context) map[string]interface{} {
	var params map[string]interface{}
	ctx.ReadJSON(&params)
	return params
}

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	DefaultOkMessage    = Message{Code: MESSAGE_OK, Msg: "ok"}
	DefaultErrorMessage = Message{Code: MESSAGE_ERROR, Msg: "error"}
)

func NewOkResponseWithData(data interface{}) Message {
	return Message{Code: MESSAGE_OK, Msg: "ok", Data: data}
}

func HandlerWeb(webController func(ctx iris.Context) api.Controller) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		// 参数绑定
		c := webController(ctx)
		cElem := reflect.ValueOf(c).Elem()
		// controller 参数校验
		if cElem.Type().Size() == 0 {
			ctx.JSON(DefaultErrorMessage)
			return
		}
		cReq := cElem.FieldByName("Req").Interface()
		if cReq == nil {
			ctx.JSON(DefaultErrorMessage)
			return
		}
		if err := c.Bind(ctx); err != nil {
			ctx.JSON(DefaultErrorMessage)
			return
		}
		if err := c.Check(ctx); err != nil {
			ctx.JSON(DefaultErrorMessage)
			return
		}
		var data interface{}
		var err error
		if data, err = c.Run(ctx); err != nil {
			ctx.JSON(DefaultErrorMessage)
			return
		}
		if data != nil {
			ctx.JSON(NewOkResponseWithData(data))
			return
		}
		ctx.JSON(DefaultOkMessage)
	}
}
