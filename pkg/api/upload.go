package api

import (
	"io/ioutil"

	"github.com/kataras/iris/v12"
	"github.com/taoey/iris-admin/pkg/entity"
	"github.com/taoey/iris-admin/pkg/service/upload"
)

func UploadAliBill(ctx iris.Context) {
	file, _, _ := ctx.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	s := string(bytes)
	upload.OnUploadAliBillPrint(s)

	result := entity.Message{
		Code: entity.MESSAGE_OK,
	}
	ctx.JSON(result)
}
