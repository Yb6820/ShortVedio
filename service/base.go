package service

import (
	"DouYin/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
example:
参数名	     位置	  类型	  必填	  说明
latest_time  query   string  否      可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
token        query   string  否      用户登录状态下设置

	{
	    "status_code": 0,
	    "status_msg": "string",
	    "next_time": 0,
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
type FeedList struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

func Getfeed(c *gin.Context) {
	token := c.Query("token")
	latest_time, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	//登录信息检验
	if token != Userbasic.Token {
		str := "登录信息失效"
		c.JSON(200, FeedList{
			StatusCode: -1,
			StatusMsg:  &str,
			NextTime:   nil,
			VideoList:  nil,
		})
		return
	}
	tm := time.Unix(latest_time, 0).Format("2006-01-02 15:04:05")
	videos := models.GetVideosBeforeTime(tm)
	video_list := make([]Video, len(videos))
	//下次传视频的最次时间nt
	var nt int64
	for k, v := range videos {
		//更新nt
		if v.CreatedAt.Unix() > nt {
			nt = v.CreatedAt.Unix()
		}
		video_list[k] = Video{
			Author:        GetUserInfoById(Userbasic.ID, v.AuthorId),
			CoverURL:      v.CoverURL,
			CommentCount:  v.CommentCount,
			FavoriteCount: v.FavoriteCount,
			ID:            int64(v.AuthorId),
			PlayURL:       v.PlayURL,
			Title:         v.Title,
			IsFavorite:    models.IsFollowOrNot(Userbasic.ID, v.AuthorId),
		}
	}
	str := "刷新视频成功!"
	rep := FeedList{
		StatusCode: 0,
		StatusMsg:  &str,
		NextTime:   &nt,
		VideoList:  video_list,
	}
	c.JSON(200, rep)
}

/*
example:
Body 参数(multipart/form-data)
参数名	类型	必填	说明
data    file   是	   视频数据
token	string 是	   用户鉴权token
title	string 是	   视频标题

	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
func PublishVideo(c *gin.Context) {

}

/*
参数名	    位置	 类型	 必填	 说明
token       query   string  是      用户鉴权token
video_id    query   string  是      视频id
action_type query   string  是      1-点赞，2-取消点赞

example:

	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
func ActFavorite(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	video := models.GetVideoById(uint(video_id))
	//登录信息检验
	if token != Userbasic.Token {
		str := "登录信息失效"
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  &str,
		})
		return
	}
	favorite := models.GetFavoriteById(uint(video_id), Userbasic.ID)
	if favorite.VideoId == 0 && action_type == 1 {
		//实现点赞总数加1,之后考虑用redis实现
		favorite.VideoId = uint(video_id)
		favorite.UserId = Userbasic.ID
		favorite.IsFavorite = true
		models.CreateFavorite(favorite)
		//更新点赞数
		models.UpdateFavoriteCount(video.FavoriteCount, true)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "点赞操作成功",
		})
	} else if action_type == 1 {
		models.UpdateFavoriteCount(video.FavoriteCount, true)
		favorite := models.Favorite{
			VideoId: uint(video_id),
			UserId:  Userbasic.ID,
		}
		models.UpdateFavorite(favorite)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "点赞操作成功",
		})
	} else {
		models.UpdateFavoriteCount(video.FavoriteCount, false)
		favorite := models.Favorite{
			VideoId: uint(video_id),
			UserId:  Userbasic.ID,
		}
		models.UpdateFavorite(favorite)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "取消点赞操作成功",
		})
	}
}

/*
example:
参数名	        位置	类型	必填	说明
token           query  string  是      用户鉴权token
video_id        query  string  是      视频id
action_type     query  string  是      1-发布评论，2-删除评论
comment_text    query  string  否      用户填写的评论内容，在action_type=1的时候使用
comment_id      query  string  否      要删除的评论id，在action_type=2的时候使用

	{
	    "status_code": 0,
	    "status_msg": "string",
	    "comment": {
	        "id": 0,
	        "user": {
	            "id": 0,
	            "name": "string",
	            "follow_count": 0,
	            "follower_count": 0,
	            "is_follow": true
	        },
	        "content": "string",
	        "create_date": "string"
	    }
	}
*/
type SendComment struct {
	Comment    *Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg"`  // 返回状态描述
}

func ActComment(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	comment_text := c.Query("comment_text")
	comment_id, _ := strconv.Atoi(c.Query("comment_id"))
	video := models.GetVideoById(uint(video_id))
	//登录信息检验
	if token != Userbasic.Token {
		str := "登录信息失效"
		c.JSON(200, SendComment{
			StatusCode: -1,
			StatusMsg:  &str,
			Comment:    nil,
		})
		return
	}
	//要返回的json格式Comment,CreateDateTime在不同功能中赋值
	com := Comment{
		Content: comment_text,
		ID:      int64(comment_id),
		User: User{
			ID:            int64(Userbasic.ID),
			FollowCount:   int64(Userbasic.Follow),
			FollowerCount: int64(Userbasic.Follower),
			Name:          Userbasic.Name,
		},
	}
	if action_type == 1 {
		comment := models.Comment{
			VideoId: uint(video_id),
			Content: comment_text,
			UserId:  Userbasic.ID,
		}
		//创建新的评论信息
		models.CreateComment(comment)
		models.UpdateCommentCount(video.CommentCount, true)
		com.CreateDate = comment.CreatedAt.String()
		str := "评论成功!"
		rep := SendComment{
			StatusCode: 0,
			StatusMsg:  &str,
			Comment:    &com,
		}
		c.JSON(200, rep)
	} else {
		comment := models.GetCommentByCommentId(uint(comment_id))
		if comment.ID != uint(comment_id) {
			str := "删除该评论失败!"
			c.JSON(200, SendComment{
				StatusCode: -1,
				StatusMsg:  &str,
				Comment:    nil,
			})
			return
		}
		str := "删除评论成功!"
		models.UpdateCommentCount(video.CommentCount, false)
		rep := SendComment{
			StatusCode: 0,
			StatusMsg:  &str,
			Comment:    &com,
		}
		c.JSON(200, rep)
	}
}

/*
example:
参数名			位置	类型	必填	说明
token      		query  string  是      用户鉴权token
to_user_id      query  string  是      对方用户id
action_type     query  string  是      1-关注，2-取消关注

	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
func ActFollow(c *gin.Context) {
	token := c.Query("token")
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	if token != Userbasic.Token {
		c.JSON(200, gin.H{
			"status_code": -1,
			"status_msg":  "登录信息失效",
		})
		return
	}
	follow := models.Follow{
		UserId:   Userbasic.ID,
		AuthorId: uint(to_user_id),
	}
	//关注某作者且关注信息不在数据库中
	if action_type == 1 && !models.IsInFollowTable(Userbasic.ID, uint(to_user_id)) {
		follow.IsFollow = true
		models.CreateFollow(follow)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "已成功关注!",
		})
		return
	} else if action_type == 1 {
		models.UpdateFollow(follow)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "已成功关注!",
		})
		return
	} else {
		models.UpdateFollow(follow)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "已取消关注!",
		})
		return
	}
}

/*
参数名			位置	类型	必填	说明
token  			query  string  是      用户鉴权token
to_user_id  	query  string  是	   对方用户id
action_type     query  string  是      1-发送消息
content			query  string  是	   消息内容

example:

	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
func SendMessage(c *gin.Context) {

}
