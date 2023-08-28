package service

import (
	"DouYin/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublishList struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

// Video
type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

/*
example:
参数名	  位置	  类型	  必填	  说明
token    query   string  是      用户鉴权token
user_id  query   string  是      用户id

	{
	    "status_code": 0,
	    "status_msg": "string",
	    "video_list": [
	        {
	            "id": 0,
	            "author": {
	                "id": 0,
	                "name": "string",
	                "follow_count": 0,
	                "follower_count": 0,
	                "is_follow": true
	            },
	            "play_url": "string",
	            "cover_url": "string",
	            "favorite_count": 0,
	            "comment_count": 0,
	            "is_favorite": true,
	            "title": "string"
	        }
	    ]
	}
*/
func GetPublishList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	usersql := models.FindUserByID(uint(user_id))
	if token != usersql.Token {
		fmt.Println("登录信息失效!")
	}
	//获取用户所有发布的视频
	publishvideo := models.GetVideosByAuthorId(uint(user_id))
	videos := make([]Video, len(publishvideo))
	for k, v := range publishvideo {
		//获取视频转换为json格式
		videos[k].ID = int64(v.ID)
		videos[k].Author = GetUserInfoById(uint(user_id), v.AuthorId)
		videos[k].CommentCount = v.CommentCount
		videos[k].PlayURL = v.PlayURL
		videos[k].CoverURL = v.CoverURL
		videos[k].FavoriteCount = v.FavoriteCount
		videos[k].Title = v.Title
		videos[k].IsFavorite = models.IsFollowOrNot(uint(user_id), uint(user_id))
	}
	str := "获取已发布的视频成功!"
	favoritelist := FavoriteList{
		StatusCode: "0",
		StatusMsg:  &str,
		VideoList:  videos,
	}
	c.JSON(200, favoritelist)
}
