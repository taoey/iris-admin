package main

import (
	"IRIS-start/src/myinit"
)

func main() {
	myinit.InitConf()
	myinit.InitMongo()
}
