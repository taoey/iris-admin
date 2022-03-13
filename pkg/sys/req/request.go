package req

import "github.com/kataras/iris/v12"

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
