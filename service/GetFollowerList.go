package service

import (
	"DouYin/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowerList struct {
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

func GetFollowerList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	usersql := models.FindUserByID(uint(user_id))
	if token != usersql.Token {
		fmt.Println("登录信息失效!")
	}
	//获取用户关注的所有作者
	followers := models.GetFollowers(uint(user_id))
	users := make([]User, len(followers))
	for k, v := range followers {
		//获取用户信息转换为json格式
		users[k] = GetUserInfoById(uint(user_id), v.UserId)
	}
	str := "获取粉丝成功"
	followerlist := FollowerList{
		StatusCode: "0",
		StatusMsg:  &str,
		UserList:   users,
	}
	c.JSON(200, followerlist)
}
