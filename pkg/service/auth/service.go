package auth

import (
	"github.com/goinggo/mapstructure"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AuthUser struct {
	Id         string   `json:"id" bson:"_id"`
	OrgId      string   `json:"orgId" bson:"orgId"`
	OrgIdPath  string   `json:"orgIdPath" bson:"orgIdPath"`
	Roles      []string `json:"roles" bson:"roles"`
	Auths      []string `json:"auths" bson:"auths"`
	OrgIds     []string `json:"orgIds" bson:"orgIds"`
	OrgIdPaths []string `json:"orgIdPaths" bson:"orgIdPaths"`
}

// session对象
var sess *sessions.Sessions

/**
 * @Description: 初始化session，时长为60分钟，需项目启动时加载
 */
func InitSession() {
	sess = sessions.New(sessions.Config{
		Cookie:  "taosid",
		Expires: time.Duration(60) * time.Minute,
		SessionIDGenerator: func() string {
			return "us:" + strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(rand.Intn(1000))
		},
	})
}

/**
 * @Description: 用户登录设置session
 * @param ctx
 * @param user 用户相关信息
 */
func AuthLogin(ctx iris.Context, user interface{}) {
	session := sess.Start(ctx)
	session.Set("authenticated", true)
	session.Set("user", user)
}

/**
 * @Description: 用户退出，注销session
 * @param ctx
 */
func AuthLogout(ctx iris.Context) {
	session := sess.Start(ctx)
	session.Delete("authenticated")
	session.Delete("user")
	sess.Destroy(ctx)
}

/**
 * @Description: 路由中间件，用于添加路由权限
 * @param auths：权限列表string，使用英文符号分割
 * @return func(ctx iris.Context)
 */
func NeedAuths(auths string) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.Values().Set("requestAuths", strings.Split(auths, ","))
		AuthRequireOne(ctx)
	}
}

/**
 * @Description: 判断用户是否有对应的相关权限
 * @param ctx
 */
func AuthRequireOne(ctx iris.Context) {
	// Check if user is authenticated
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	requestAuths := ctx.Values().GetDefault("requestAuths", []string{}).([]string)             //路由中的权限
	userAuths := sess.Start(ctx).Get("user").(map[string]interface{})["auths"].([]interface{}) //用户拥有的权限

	hasAuth := false
	for i := 0; i < len(userAuths); i++ {
		if hasAuth || userAuths[i] == "AUTH_ROOT" {
			hasAuth = true
			break
		} else {
			for j := 0; j < len(requestAuths); j++ {
				if userAuths[i] == requestAuths[j] {
					hasAuth = true
					break
				}
			}
		}
	}

	if hasAuth {
		// 更新redis对应授权key的TTL
		sess.ShiftExpiration(ctx)
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}
}

/**
 * @Description: 用户登录状态查询
 * @param ctx
 * @return value
 */
func AuthCurrentUserGetEx(ctx iris.Context) (value map[string]interface{}) {
	sessUser := sess.Start(ctx).Get("user")
	if sessUser == nil {
		value = nil
	} else {
		mapstructure.Decode(sessUser, &value)
	}
	return
}

/**
 * @Description: 获取当前登录用户
 * @param ctx
 * @return *AuthUser
 */
func AuthCurrentUserGet(ctx iris.Context) *AuthUser {
	sessUser := sess.Start(ctx).Get("user")
	if sessUser == nil {
		return nil
	} else {
		authUser := AuthUser{}
		mapstructure.Decode(sessUser, &authUser)
		return &authUser
	}
}

/**
 * @Description:  权限检查
 * @param requireAuth
 * @param userAuths
 * @return bool
 */
func AuthCheck(requireAuth string, userAuths []string) bool {
	for i := 0; i < len(userAuths); i++ {
		if userAuths[i] == "AUTH_ROOT" || userAuths[i] == requireAuth {
			return true
		}
	}
	return false
}
