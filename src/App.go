package main

import (
	"IRIS-start/src/myinit"
	"github.com/kataras/iris"
	"log"
)

func main() {
	// 初始化
	myinit.InitConf()
	//myinit.InitMongo()
	//myinit.InitQuartz()

	// 初始化App
	App = iris.New()
	SetRoutes()

	// 启动
	run := App.Run(iris.Addr(myinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}
