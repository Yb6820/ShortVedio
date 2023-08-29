package service

import (
	"DouYin/middleware"
	"DouYin/models"
	"DouYin/utils"
	"DouYin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToRegister(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("username")
	password := c.Query("password")

	if user.Name == "" || password == "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "用户名或密码不能为空",
			"user_id":     0,
			"token":       "",
		})
		return
	}
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "该用户名已存在",
			"user_id":     0,
			"token":       "",
		})
		return
	}
	user.Password = utils.MakePassword(password)
	token, tErr := middleware.SetToken(user.Name)
	err := models.CreateUser(user)
	if err != nil || tErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "新增用户失败",
			"user_id":     0,
			"token":       "",
		})
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "新增用户成功",
		"user_id":     user.ID,
		"token":       token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := models.FindUserByName(username)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "该用户不存在",
			"user_id":     user.ID,
			"token":       "",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Password)
	//fmt.Println("user.Salt:", user.Salt, "password:", password, "User.Password", user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "密码不正确",
			"user_id":     user.ID,
			"token":       "",
		})
		return
	}
	//生成token
	token, err := middleware.SetToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": errmsg.ERROR,
			"status_msg":  errmsg.GetErrMsg(errmsg.ERROR),
			"user_id":     0,
			"token":       "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "登陆成功",
		"user_id":     user.ID,
		"token":       token,
	})
	return
}
