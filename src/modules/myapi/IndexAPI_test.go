package myapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIndexHelloJson(t *testing.T) {
	s := Student{
		Id:   12,
		Name: "tao",
	}
	bytes, _ := json.Marshal(s)
	fmt.Println(string(bytes))
}

func Test02(t *testing.T) {
	user := User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}
