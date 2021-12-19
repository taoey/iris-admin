package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestHttpGet(t *testing.T) {
	for i := 0; i < 10; i++ {
		parm := map[string]string{
			"key": "",
			"ip":  "" + strconv.Itoa(i),
		}
		s, err := httpGet("", parm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}

}

func TestHttpPost(t *testing.T) {
	var test = test{Obj: "{\"aa\":\"aa\"}"}
	info, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	s, err := HTTPPost("", "application/json", string(info))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

func TestHttpPostForm(t *testing.T) {
	//	HttpRequest()
}

type test struct {
	Obj string `json:"obj"`
}
