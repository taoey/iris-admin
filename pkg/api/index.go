package api

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
)

// route: "/"

func Index(ctx iris.Context) {
	ctx.WriteString("hello world -- from isis")
}

func IndexHelloJson(ctx iris.Context) {
	s := Student{
		Id:   12,
		Name: "tao",
	}
	bytes, _ := json.Marshal(s)
	fmt.Println(string(bytes))
	ctx.JSON(s)
}

// 变量名称必须大写否则转化json的时候不能进行转化
type Student struct {
	Id   int
	Name string
}

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}
