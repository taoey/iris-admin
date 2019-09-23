package myinit

import (
	"github.com/Taoey/iris-cli/src/system/job"
	"github.com/bamzi/jobrunner"
)

func InitQuartz() {
	jobrunner.Start()
	jobrunner.Schedule("@every 2s", job.PrintTime{})
}
