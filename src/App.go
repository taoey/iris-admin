package main

import (
	"IRIS-start/src/myinit"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"reflect"
)

func main() {
	myinit.InitConf()
	//myinit.InitMongo()
	myinit.InitQuartz()
	// 启动
	app := iris.New()
	//根路由
	app.Get("/", func(context iris.Context) {
		context.WriteString("hello world -- from isis")
	})

	app.Post("/upload/wx_bill.json", iris.LimitRequestBodySize(5<<20), func(ctx iris.Context) {
		file, _, _ := ctx.FormFile("file")
		bytes, _ := ioutil.ReadAll(file)
		s := string(bytes)
		fmt.Println(reflect.TypeOf(bytes), s)
	})

	run := app.Run(iris.Addr(myinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	fmt.Println(run)
}
