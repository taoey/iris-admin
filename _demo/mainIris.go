package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	htmlEngin := iris.HTML("iris/view/", ".html")
	app.RegisterView(htmlEngin)

	//根路由
	app.Get("/", func(context iris.Context) {
		context.WriteString("hello world -- from isis")
	})

	app.Get("/print_json_1", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "hello json",
		})
	})
	app.Get("/print_json_2", func(ctx iris.Context) {
		myMap := make(map[string]interface{})
		myMap["code"] = 200
		myMap["data"] = map[string]string{
			"name": "tao",
			"age":  "hello",
		}
		ctx.JSON(myMap)
	})

	app.Get("/hello", func(context iris.Context) {
		context.ViewData("title", "MyFirstIrisDemo")
		context.ViewData("content", "hello iris ")
		context.View("hello.html")
	})

	run := app.Run(iris.Addr("127.0.0.1:8080"), iris.WithCharset("UTF-8"))
	fmt.Println(run)
}
