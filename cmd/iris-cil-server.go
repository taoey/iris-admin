package main

import (
	"fmt"
	"github.com/Taoey/iris-cli/src/pkg/api"
	"github.com/Taoey/iris-cli/src/sysinit"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris"
	"log"
)

var App *iris.Application
var HttpLimithandler iris.Handler

//程序入口
func main() {
	// 初始化
	sysinit.InitConf()
	//sysinit.InitMongo()
	//sysinit.InitQuartz()
	sysinit.InitMysql()

	// 初始化App
	App = iris.New()

	SetHttpLimiter()
	SetRoutes()

	// 启动
	run := App.Run(iris.Addr(sysinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}

// 设置路由
func SetRoutes() {

	//主页
	App.Get("/", api.Index)
	App.Get("/hello_json", HttpLimithandler, api.IndexHelloJson)

	//根API
	RootApi := App.Party("api/v1")

	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), api.UploadAliBill)

	//download
	RootApi.Get("/download/demo1", api.ApiDownloadDemo1)
	RootApi.Get("/download/demo2", api.ApiDownloadDemo2)
	RootApi.Get("/download/demo3", api.ApiDownloadDemo3)
	RootApi.Get("/download/demo4", api.ApiDownloadLimite)
}

// 设置HTTP限流器
func SetHttpLimiter() {
	// 每秒1个请求
	limiter := tollbooth.NewLimiter(1, nil)
	HttpLimithandler = LimitHandler(limiter)
}

// 引用自：https://github.com/didip/tollbooth_iris/blob/master/tollbooth_iris.go 但是原项目没有人维护了，其对应的版本为iris v6
// 因此我们把对应的代码下载，自己进行维护
// LimitHandler is a middleware that performs
// rate-limiting given a "limiter" configuration.
func LimitHandler(lmt *limiter.Limiter) iris.Handler {
	return func(ctx iris.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.ResponseWriter(), ctx.Request())
		fmt.Println(httpError)
		if httpError != nil {
			ctx.StatusCode(httpError.StatusCode)
			ctx.WriteString(httpError.Message)
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}
