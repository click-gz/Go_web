package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Tel  string `gorm:"type:varchar(11);not null;unique"`
	Pwd  string `gorm:"size:255;not null"`
}

//数据库配置
const (
	userName = "root"
	password = "pigz"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "learn"
)

func InitDB() *gorm.DB {

	db, err := gorm.Open("mysql", strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, ""))
	if err != nil {
		panic("failed to connect database: " + err.Error())

	}
	db.AutoMigrate(&User{})

	return db

}

func CheckTel(db *gorm.DB, tel string) bool {
	var user User
	db.Where("Tel = ?", tel).First(&user)
	fmt.Println(user.ID)
	if user.ID != 0 {
		return true
	}
	return false
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.GET("api/user/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		tel := c.PostForm("tel")
		pwd := c.PostForm("pwd")

		//数据验证
		if len(tel) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
			return
		}
		if len(pwd) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于6位",
			})
			return
		}
		fmt.Println(name, " ", tel, " ", pwd)

		//判断手机号是否存在
		if CheckTel(db, tel) { //用户存在
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已存在",
			})
			return
		}

		//创建用户
		newUser := User{
			Name: name,
			Tel:  tel,
			Pwd:  pwd,
		}
		db.Create(&newUser)

		//返回结果

		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
