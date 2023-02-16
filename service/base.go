package service

import "github.com/gin-gonic/gin"

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
	c.JSON(200, gin.H{})
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
