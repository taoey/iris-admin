package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris/v12"
)

var HttpLimithandler iris.Handler

// 设置HTTP限流器
func InitHttpLimiter() {
	// 每秒1个请求
	limiter := tollbooth.NewLimiter(1, nil)
	HttpLimithandler = limitHandler(limiter)
}

// 引用自：https://github.com/didip/tollbooth_iris/blob/master/tollbooth_iris.go 但是原项目没有人维护了，其对应的版本为iris v6
// 因此我们把对应的代码下载，自己进行维护
// LimitHandler is a middleware that performs
// rate-limiting given a "limiter" configuration.
func limitHandler(lmt *limiter.Limiter) iris.Handler {
	return func(ctx iris.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.ResponseWriter(), ctx.Request())
		if httpError != nil {
			ctx.StatusCode(httpError.StatusCode)
			ctx.WriteString(httpError.Message)
			ctx.StopExecution()
			return
		}
		ctx.Next()
	}
}
