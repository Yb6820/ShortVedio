package router

import (
	"DouYin/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//视频流接口
	r.GET("/douyin/feed", service.Getfeed)

	//用户登录和注册模块
	r.POST("/douyin/user/register/", service.ToRegister)
	r.POST("/douyin/user/login/", service.Login)
	r.GET("/douyin/user/", service.GetUserInfo)
	return r
}
