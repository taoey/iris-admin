package httpclient

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestHttpGet(t *testing.T) {
	for i := 0; i < 10; i++ {
		parm := map[string]string{
			"key": "417dc94ff687da4b1bc8fa89e7353445",
			"ip":  "124.72.37.20" + strconv.Itoa(i),
		}
		s, err := httpGet("http://restapi.amap.com/v3/ip", parm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}

}

func TestHttpPost(t *testing.T) {
	var test = test{Obj: "{\"getwayMac\":\"39FFD505474D383737780643\"}"}
	info, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	s, err := HttpPost("http://124.193.136.53:6200//LedService.svc/GetDeviceGateWayDetail", "application/json", string(info))
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
