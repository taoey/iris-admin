package main

import (
	"github.com/Taoey/iris-cli/src/pkg/api"
	"github.com/Taoey/iris-cli/src/sysinit"
	"github.com/kataras/iris"
	"log"
)

var App *iris.Application

//程序入口
func main() {
	// 初始化
	sysinit.InitConf()
	//myinit.InitMongo()
	//myinit.InitQuartz()

	// 初始化App
	App = iris.New()
	SetRoutes()

	// 启动
	run := App.Run(iris.Addr(sysinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}

// 设置路由
func SetRoutes() {

	//主页
	App.Get("/", api.Index)
	App.Get("/hello_json", api.IndexHelloJson)

	//根API
	RootApi := App.Party("api/")

	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), api.UploadAliBill)

}
