package sysinit

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"time"
)

func InitQuartz() {
	jobrunner.Start()
	jobrunner.Schedule("@every 2s", PrintTime{})
}

//打印时间任务
type PrintTime struct {
}

func (p PrintTime) Run() {
	fmt.Println("time:", time.Now())
}
