package router

import (
	"DouYin/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//视频流接口
	r.GET("/douyin/feed", service.Getfeed)

	//投稿接口
	r.POST("/douyin/publish/action/", service.PublishVideo)

	//点赞操作
	r.POST("/douyin/favorite/action/", service.ActFavorite)

	//评论操作
	r.POST("/douyin/comment/action/", service.ActComment)

	//关注操作
	r.POST("/douyin/relation/action/", service.ActFollow)

	//发送消息
	r.POST("/douyin/message/action/", service.SendMessage)

	//用户登录和注册模块
	r.POST("/douyin/user/register/", service.ToRegister)
	r.POST("/douyin/user/login/", service.Login)
	//获取用户信息
	r.GET("/douyin/user/", service.GetUserInfo)

	//发布视频列表
	r.GET("/douyin/publish/list/", service.GetPublishList)
	//喜欢列表
	r.GET("/douyin/favorite/list/", service.GetFavoriteList)
	//评论列表
	r.GET("/douyin/comment/list/", service.GetCommentList)
	//关注列表
	r.GET("/douyin/relation/follow/list/", service.GetFollowList)
	//粉丝列表
	r.GET("/douyin/relation/follower/list/", service.GetFollowerList)
	//好友列表
	r.GET("/douyin/relation/friend/list/", service.GetFriendList)
	//获取聊天记录
	r.GET("/douyin/message/chat/", service.GetChatHistory)
	return r
}
