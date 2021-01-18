package user

import (
	. "github.com/Taoey/iris-cli/pkg"
	"github.com/Taoey/iris-cli/pkg/service/auth"
	"github.com/Taoey/iris-cli/pkg/sys/req"
	"github.com/Taoey/iris-cli/pkg/sys/resp"
	"github.com/kataras/iris"
)

// 用户登录
func UserLoginHandler(ctx iris.Context) {
	params := req.Read2Map(ctx)
	auth.AuthLogin(ctx, params)
	LOG.Debug("user login:", params)
	ctx.JSON(resp.OkResponse())
}

// 查看系统当前用户
func UserCurrentHandler(ctx iris.Context) {
	user := auth.AuthCurrentUserGetEx(ctx)
	ctx.JSON(resp.OkResponseWithRet(user))
}

// 用户登出
func UserLogoutHandler(ctx iris.Context) {
	user := auth.AuthCurrentUserGetEx(ctx)
	LOG.Debug("user logout:", user)
	auth.AuthLogout(ctx)
	ctx.JSON(resp.OkResponse())
}
