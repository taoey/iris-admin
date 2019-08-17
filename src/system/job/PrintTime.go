package job

import (
	"fmt"
	"time"
)

//打印时间任务
type PrintTime struct {
}

func (p PrintTime) Run() {
	fmt.Println("time:", time.Now())
}
