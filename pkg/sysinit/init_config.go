package sysinit

import (
	"github.com/olebedev/config"
	"os"
)

var GCF *config.Config //global config

func InitConf() {
	pwd, _ := os.Getwd()
	configPath := pwd + "/configs/application.yml"
	GCF, _ = config.ParseYamlFile(configPath)
}
