package sysinit

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// mysql 连接（已经配置了连接池）
var MysqlSession *gorm.DB

// mysql 初始化
func InitMysql() {
	db, err := gorm.Open("mysql", GCF.UString("mysql.url"))
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(GCF.UInt("mysql.max_idle", 10))  //空闲最大连接数
	db.DB().SetMaxOpenConns(GCF.UInt("mysql.max_open", 100)) //最大连接数

	db.LogMode(true)
	//db.SetLogger(Logger{})
	MysqlSession = db

	//fmt.Println("mysql init end",db)
}
