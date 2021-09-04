package UserControl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go_v/common"
	"go_v/model"
	"net/http"
)

func CheckTel(db *gorm.DB, tel string) bool {
	var user user.User
	db.Where("Tel = ?", tel).First(&user)
	fmt.Println(user.ID)
	if user.ID != 0 {
		return true
	}
	return false
}
func CheckPwd(db *gorm.DB, tel string, pwd string) bool {
	var user user.User
	db.Where("Tel = ?", tel).First(&user)
	fmt.Println(user.Pwd, " ", pwd)
	if user.Pwd == pwd {
		return true
	}
	return false
}

func Register(c *gin.Context) {
	db := common.GetDB()

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
	newUser := user.User{
		Name: name,
		Tel:  tel,
		Pwd:  pwd,
	}
	db.Create(&newUser)

	//返回结果

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()

	//获取参数
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

	//判断手机号是否存在
	if !CheckTel(db, tel) { //用户不存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	if !CheckPwd(db, tel, pwd) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
