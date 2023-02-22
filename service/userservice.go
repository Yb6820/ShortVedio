package service

import (
	"DouYin/models"
	"DouYin/utils"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func ToRegister(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("username")
	password := c.Query("password")

	//获取一个生成密码的随机数
	salt := fmt.Sprintf("%06d", rand.Int31())
	if user.Name == "" || password == "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "用户名或密码不能为空",
			"user_id":     0,
			"token":       "string",
		})
		return
	}
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "该用户名已存在",
			"user_id":     0,
			"token":       "string",
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "新增用户成功",
		"user_id":     user.ID,
		"token":       "string",
	})
}

func Login(c *gin.Context) {
	/* name := c.Query("name")
	password := c.Query("password") */
	username := c.Query("username")
	password := c.Query("password")
	user := models.FindUserByName(username)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "该用户不存在",
			"user_id":     user.ID,
			"token":       "string",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	//fmt.Println("user.Salt:", user.Salt, "password:", password, "User.Password", user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "密码不正确",
			"user_id":     user.ID,
			"token":       "string",
		})
		return
	}
	//校验密码
	pwd := utils.MakePassword(password, user.Salt)
	data := models.FindUserByNameAndPwd(username, pwd)
	Userbasic = user
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "登陆成功",
		"user_id":     data.ID,
		"token":       data.Token,
	})
}
