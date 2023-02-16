package service

import "github.com/gin-gonic/gin"

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

}
