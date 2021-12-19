package router

import (
	"github.com/kataras/iris"
	"github.com/taoey/iris-admin/pkg/api"
	"github.com/taoey/iris-admin/pkg/service/test"
	"github.com/taoey/iris-admin/router/middleware"
)

// 设置路由
func SetRoutes(app *iris.Application) {
	//主页
	app.Get("/", api.Index)
	app.Get("/hello_json", middleware.HttpLimithandler, api.IndexHelloJson)

	//根API
	rootApi := app.Party("api/v1")

	// upload
	rootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), api.UploadAliBill)

	//download
	rootApi.Get("/download/demo1", api.ApiDownloadDemo1)
	rootApi.Get("/download/demo2", api.ApiDownloadDemo2)
	rootApi.Get("/download/demo3", api.ApiDownloadDemo3)
	rootApi.Get("/download/demo4", api.ApiDownloadLimite)
	rootApi.Get("/download/demo5", api.ApiDownloadLimiteSleep)
	rootApi.Get("/download/demo6", api.ApiDownloadDemo6)
	rootApi.Post("/test/map_parms", test.MapParmsHandler)

	// 用户登录登出
	// rootApi.Post("/user/login", user.UserLoginHandler)
	// rootApi.Get("/user/current", user.UserCurrentHandler)
	// rootApi.Get("/user/logout", user.UserLogoutHandler)

	// 测试
	rootApi.Get("/test/error/zero", test.ErrorHandler)
}
