package main

import (
	"IRIS-start/src/modules/myapi"
	"github.com/kataras/iris"
)

// 设置路由
func SetRoutes() {

	//主页
	App.Get("/", myapi.Index)
	App.Get("/hello_json", myapi.IndexHelloJson)

	//根API
	RootApi := App.Party("api/")

	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), myapi.UploadAliBill)

}
