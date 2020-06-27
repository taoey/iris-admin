package api

import (
	. "github.com/Taoey/iris-cli/src/entity"
	"github.com/Taoey/iris-cli/src/pkg/service"
	"github.com/kataras/iris"
	"io/ioutil"
)

func UploadAliBill(ctx iris.Context) {
	file, _, _ := ctx.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	s := string(bytes)
	service.OnUploadAliBillPrint(s)

	result := Message{
		Code: MESSAGE_OK,
	}
	ctx.JSON(result)
}
