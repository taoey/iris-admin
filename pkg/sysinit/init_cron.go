package sysinit

import (
	"fmt"
	"time"

	"github.com/bamzi/jobrunner"
)

func InitCron() {
	jobrunner.Start()
	// jobrunner.Schedule("@every 2s", PrintTime{})

}

// 打印时间任务
type PrintTime struct {
}

func (p PrintTime) Run() {
	fmt.Println("time:", time.Now())
}

// docker run -dit -p 8797:9090 -v /root/soft/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
