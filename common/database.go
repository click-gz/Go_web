package common

import (
	"github.com/jinzhu/gorm"
	"go_v/model"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "pigz"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "learn"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open("mysql", strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, ""))
	if err != nil {
		panic("failed to connect database: " + err.Error())

	}
	db.AutoMigrate(&user.User{})
	DB = db

}

func GetDB() *gorm.DB {
	return DB
}
