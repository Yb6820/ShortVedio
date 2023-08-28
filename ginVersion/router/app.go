package router

import (
	"DouYin/middleware"
	"DouYin/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//投稿接口
		auth.POST("/douyin/publish/action/", service.PublishVideo)
		//点赞操作
		auth.POST("/douyin/favorite/action/", service.ActFavorite)
		//评论操作
		auth.POST("/douyin/comment/action/", service.ActComment)
		//关注操作
		auth.POST("/douyin/relation/action/", service.ActFollow)
		//发送消息
		auth.POST("/douyin/message/action/", service.SendMessage)
		//获取用户信息
		auth.GET("/douyin/user/", service.GetUserInfo)
		//发布视频列表
		auth.GET("/douyin/publish/list/", service.GetPublishList)
		//喜欢列表
		auth.GET("/douyin/favorite/list/", service.GetFavoriteList)
		//评论列表
		auth.GET("/douyin/comment/list/", service.GetCommentList)
		//关注列表
		auth.GET("/douyin/relation/follow/list/", service.GetFollowList)
		//粉丝列表
		auth.GET("/douyin/relation/follower/list/", service.GetFollowerList)
		//好友列表
		auth.GET("/douyin/relation/friend/list/", service.GetFriendList)
		//获取聊天记录
		auth.GET("/douyin/message/chat/", service.GetChatHistory)
	}
	normal := r.Group("api/v1")
	{
		//视频流接口
		normal.GET("/douyin/feed/", service.Getfeed)
		//用户登录和注册模块
		normal.POST("/douyin/user/register/", service.ToRegister)
		normal.POST("/douyin/user/login/", service.Login)
	}
	return r
}
