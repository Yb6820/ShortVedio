package service

import "github.com/gin-gonic/gin"

type FriendList struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户列表
}

/*
example:
参数名	  位置	  类型	  必填	  说明
user_id   query  string  是      用户id
token     query  string  是      用户鉴权token

{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true
        }
    ]
}
*/
func GetFriendList(c *gin.Context) {

}
