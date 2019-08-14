package myinit

import (
	"github.com/olebedev/config"
	"os"
	"strings"
)

var GCF *config.Config //global config

func InitConf() {
	pwd, _ := os.Getwd()
	pwd = strings.Replace(pwd, "src\\myinit", "", -1) // 使用的替换方法，所以项目路径中不能包括init字符
	configPath := pwd + "/src/config/application.yml"
	GCF, _ = config.ParseYamlFile(configPath)
}
