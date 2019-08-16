package myinit

import (
	"IRIS-start/src/system/job"
	"github.com/bamzi/jobrunner"
)

func InitQuartz() {
	jobrunner.Start()
	jobrunner.Schedule("@every 2s", job.PrintTime{})
}
