package myapi

import (
	. "IRIS-start/src/entity"
	"IRIS-start/src/modules/myservice"
	"github.com/kataras/iris"
	"io/ioutil"
)

func UploadAliBill(ctx iris.Context) {
	file, _, _ := ctx.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	s := string(bytes)
	myservice.UploadAliBillPrint(s)

	result := Message{
		Code: MESSAGE_OK,
	}
	ctx.JSON(result)
}
