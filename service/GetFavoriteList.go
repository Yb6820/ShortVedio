package service

import (
	"DouYin/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteList struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}

/*
example:
参数名	  位置	  类型	  必填	  说明
user_id   query  string  是      用户id
token     query  string  是      用户鉴权token


{
    "status_code": "string",
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

func GetFavoriteList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	usersql := models.FindUserByID(uint(user_id))
	if token != usersql.Token {
		fmt.Println("登录信息失效!")
	}
	//获取用户所有点赞的视频
	favoritevideo := models.GetFavoriteListByUserId(uint(user_id))
	videos := make([]Video, len(favoritevideo))
	for k, v := range favoritevideo {
		//获取视频转换为json格式
		video := models.GetVideoById(v.VideoId)
		user := models.FindUserByID(video.AuthorId)
		videos[k].ID = int64(video.ID)
		videos[k].Author = User{
			ID:            int64(user.ID),
			FollowCount:   int64(user.Follow),
			FollowerCount: int64(user.Follower),
			Name:          user.Name,
			IsFollow:      models.IsFollowOrNot(Userbasic.ID, user.ID),
		}
		videos[k].CommentCount = video.CommentCount
		videos[k].PlayURL = video.PlayURL
		videos[k].CoverURL = video.CoverURL
		videos[k].FavoriteCount = video.FavoriteCount
		videos[k].Title = video.Title
		videos[k].IsFavorite = true
	}
	str := "获取已点赞的视频成功"
	favoritelist := FavoriteList{
		StatusCode: "0",
		StatusMsg:  &str,
		VideoList:  videos,
	}
	c.JSON(200, favoritelist)
}
