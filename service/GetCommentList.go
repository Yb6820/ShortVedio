package service

import (
	"DouYin/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentList struct {
	CommentList []Comment `json:"comment_list"` // 评论列表
	StatusCode  int64     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

// Comment
type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

/*
example:
参数名	   位置	   类型	   必填	   说明
token     query   string  是      用户鉴权token
video_id  query   string  是      视频id

	{
	    "status_code": 0,
	    "status_msg": "string",
	    "comment_list": [
	        {
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
	    ]
	}
*/
func GetCommentList(c *gin.Context) {
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	token := c.Query("token")
	if token != Userbasic.Token {
		fmt.Println("登录信息失效!")
		return
	}
	//获取视频底下的所有评论
	comments := models.GetCommentByVideoId(uint(video_id))
	commentlist := make([]Comment, len(comments))
	for k, v := range comments {
		user := GetUserInfoById(v.UserId, models.GetVideoById(v.VideoId).ID)
		commentlist[k].ID = int64(v.ID)
		commentlist[k].User = user
		commentlist[k].Content = v.Content
		commentlist[k].CreateDate = v.CreatedAt.Format("01-02")
	}
	str := "获取所有评论成功"
	res := CommentList{
		StatusCode:  0,
		StatusMsg:   &str,
		CommentList: commentlist,
	}
	c.JSON(200, res)
}
