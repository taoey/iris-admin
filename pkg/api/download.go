package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/juju/ratelimit"
	"github.com/kataras/iris/v12"
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
	defer func() {
		fmt.Println("download limit closed")
	}()
	// 断点续传
	request := ctx.Request()
	var start, end int64
	fmt.Sscanf(request.Header.Get("Range"), "bytes=%d-%d", &start, &end)
	fmt.Println(start, end)

	//filedir :="./files/"
	//filename := "android-studio-ide-191.5791312-windows.exe"
	//filepath := filedir + filename
	filename := "data.txt"
	//filepath := "./files/data.txt"
	filepath := `G:\softdata\os\ubuntu-18.04.3-desktop-amd64.iso`

	//var downloadSpeed float64 = 1024 * 1024 * 1 //下载速度 1MB/s
	var downloadSpeed float64 = 100 * 1 * 1
	var takeTokenCount int64 = 100
	var timeOut int64 = 1000 * 60 // 设置超时时间 60s

	f, _ := os.Open(filepath)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	current := start
	bucket := ratelimit.NewBucketWithRate(downloadSpeed, 10000)

	// 下载监控
	//go func() {
	//	for {
	//		fmt.Println(bucket.Available(), current, len(data), current/len(data)*100)
	//		time.Sleep(time.Millisecond * 200)
	//	}
	//}()

	startTime := time.Now().UnixNano() / 1e6
	for current < end && current < int64(len(data)) {
		currentTime := time.Now().UnixNano() / 1e6
		if currentTime-startTime <= timeOut {
			bucket.Wait(takeTokenCount)
			//TODO 该处下载有问题可能会造成异常 目测cap容量比len容量大很多，不必担心
			//fmt.Println(len(data),cap(data),current,current+int64(takeTokenCount),data[current : current+int64(takeTokenCount)])
			ctx.ResponseWriter().Write(data[current : current+int64(takeTokenCount)])
			current = current + int64(takeTokenCount)
		} else {
			ctx.ResponseWriter().EndResponse()
			return
		}
	}
}

// 使用sleep限速下载 : 测试不通过
func ApiDownloadLimiteSleep(ctx iris.Context) {

	//filedir :="./files/"
	//filename := "android-studio-ide-191.5791312-windows.exe"
	//filepath := filedir + filename
	filename := "data.txt"
	filepath := "./files/data.txt"

	//var downloadSpeed float64 = 1024 * 1024 * 1 //下载速度 1MB/s
	var takeTokenCount = 100
	var timeOut int64 = 1000 * 30 // 设置超时时间 60s

	f, _ := os.Open(filepath)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// 下载监控
	//go func() {
	//	for {
	//		fmt.Println(bucket.Available(), current, len(data), current/len(data)*100)
	//		time.Sleep(time.Millisecond * 200)
	//	}
	//}()

	current := 0
	startTime := time.Now().UnixNano() / 1e6
	for current < len(data) {
		time.Sleep(time.Second * 1)
		currentTime := time.Now().UnixNano() / 1e6
		if currentTime-startTime <= timeOut {
			//TODO 该处下载有问题可能会造成异常 目测cap容量比len容量大很多，不必担心
			fmt.Println(len(data), cap(data), current, current+takeTokenCount, data[current:current+takeTokenCount])
			ctx.ResponseWriter().Write(data[current : current+takeTokenCount])
			current = current + takeTokenCount
		} else {
			ctx.ResponseWriter().EndResponse()
			return
		}
	}
}

func ApiDownloadDemo6(ctx iris.Context) {
	defer func() {
		fmt.Println("download6 closed")
	}()
	filepath := `G:\softdata\os\ubuntu-18.04.3-desktop-amd64.iso`
	ctx.SendFile(filepath, "ubuntu-18.04.3-desktop-amd64.iso")
}
