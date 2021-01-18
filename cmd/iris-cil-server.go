package main

import (
	"fmt"
	"github.com/Taoey/iris-cli/pkg/api"
	"github.com/Taoey/iris-cli/pkg/service/auth"
	"github.com/Taoey/iris-cli/pkg/service/test"
	"github.com/Taoey/iris-cli/pkg/service/user"
	"github.com/Taoey/iris-cli/pkg/sysinit"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"log"
)

var App *iris.Application
var HttpLimithandler iris.Handler

//程序入口
func main() {
	// 初始化App
	App = iris.New()

	// 初始化
	sysinit.InitConf()
	SetMiddleware()
	sysinit.InitLogger()
	auth.InitSession()

	//sysinit.InitMongo()
	//sysinit.InitQuartz()
	//sysinit.InitMysql()

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
	RootApi.Get("/download/demo5", api.ApiDownloadLimiteSleep)
	RootApi.Get("/download/demo6", api.ApiDownloadDemo6)
	RootApi.Post("/test/map_parms", test.MapParmsHandler)

	// 用户登录登出
	RootApi.Post("/user/login", user.UserLoginHandler)
	RootApi.Get("/user/current", user.UserCurrentHandler)
	RootApi.Get("/user/logout", user.UserLogoutHandler)

	// 测试
	RootApi.Get("/test/error/zero", test.ErrorHandler)
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

func SetMiddleware() {
	// 设置未知异常捕获
	App.Use(recover.New())
	// 设置限流器
	SetHttpLimiter()
}
