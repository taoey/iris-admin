package api

import (
	"fmt"
	"github.com/juju/ratelimit"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
	"time"
)

// 官方下载示例 https://www.studyiris.com/example/fileServer/sendFiles.html
// 推荐使用官方的代码进行下载，其实如果看源码的话，源码使用的下载模式和第二种相同
func ApiDownloadDemo1(ctx iris.Context) {
	pwd, _ := os.Getwd()

	filedir := pwd + "/files/"
	filename := "data1.txt"
	filepath := filedir + filename

	ctx.SendFile(filepath, filename)
}

// 互联网示例：go实现上传和下载excel接口 https://blog.csdn.net/weixin_43456598/article/details/100696033

// 通过自己设置header的方式下载
func ApiDownloadDemo2(ctx iris.Context) {
	pwd, _ := os.Getwd()

	filedir := pwd + "/files/"
	filename := "data1.txt"
	filepath := filedir + filename

	f, _ := os.Open(filepath)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Write(data)
}

// 使用相对路径进行下载
func ApiDownloadDemo3(ctx iris.Context) {
	filename := "./data1.txt1"
	filepath := filename

	ctx.SendFile(filepath, "11.txt")
}

// url对应的资源需要配置Nginx服务器
func SendURLFile(ctx iris.Context) {
	ctx.Redirect("http://url", 302)
}

// 使用令牌桶限速下载
func ApiDownloadLimite(ctx iris.Context) {
	//
	//filedir :="./files/"
	//filename := "android-studio-ide-191.5791312-windows.exe"
	//filepath := filedir + filename
	filename := "getfusion_temp.dmg"
	filepath := "/Users/tao/Downloads/getfusion.dmg"

	var downloadSpeed float64 = 1024 * 1024 * 1
	var takeTokenCount int64 = 200

	f, _ := os.Open(filepath)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	current := 0
	bucket := ratelimit.NewBucketWithRate(downloadSpeed, 10000) //下载速度 1MB/s

	// 下载监控
	go func() {
		for {
			fmt.Println(bucket.Available(), current, len(data), current/len(data)*100)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	var timeOut int64 = 1000 * 60 // 设置超时时间 60s

	startTime := time.Now().UnixNano() / 1e6
	for current < len(data) {
		currentTime := time.Now().UnixNano() / 1e6
		if currentTime-startTime <= timeOut {
			bucket.Wait(takeTokenCount)
			ctx.ResponseWriter().Write(data[current : current+int(takeTokenCount)])
			current = current + int(takeTokenCount)
		} else {
			ctx.ResponseWriter().CloseNotify()
			return
		}
	}
}
