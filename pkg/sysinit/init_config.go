package sysinit

import (
	"os"

	"github.com/olebedev/config"
)

var GCF *config.Config //global config

func InitConf() {
	pwd, _ := os.Getwd()
	configPath := pwd + "/configs/application.yml"
	GCF, _ = config.ParseYamlFile(configPath)
}
