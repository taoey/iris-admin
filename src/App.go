package main

import (
	"github.com/kataras/iris"
	"log"
)

func main() {
	app := iris.New()
	//根路由
	app.Get("/", func(context iris.Context) {
		context.WriteString("hello world -- from isis")
	})

	run := app.Run(iris.Addr("127.0.0.1:8080"), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}
