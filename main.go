package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/taoey/iris-admin/pkg/sysinit"
	"github.com/taoey/iris-admin/router"
	"github.com/taoey/iris-admin/router/middleware"
)

var App *iris.Application

//程序入口
func main() {
	// 初始化App
	App = iris.New()

	// 初始化
	sysinit.InitConf()
	InitMiddleware()
	sysinit.InitLogger()

	//sysinit.InitMongo()
	sysinit.InitCron()
	//sysinit.InitMysql()

	router.SetRoutes(App)

	// 启动
	run := App.Run(iris.Addr(sysinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}

func InitMiddleware() {
	// 设置未知异常捕获
	App.Use(recover.New())
	// 设置限流器
	middleware.InitHttpLimiter()
}
