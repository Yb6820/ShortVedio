package service

import (
	"DouYin/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowList struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户信息列表
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
func GetFollowList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	usersql := models.FindUserByID(uint(user_id))
	if token != usersql.Token {
		fmt.Println("登录信息失效!")
	}
	//获取用户关注的所有作者
	follows := models.GetFollows(uint(user_id))
	users := make([]User, len(follows))
	for k, v := range follows {
		//获取用户信息转换为json格式
		user := models.FindUserByID(v.AuthorId)
		users[k].ID = int64(user.ID)
		users[k].FollowCount = int64(user.Follow)
		users[k].FollowerCount = int64(user.Follower)
		users[k].Name = user.Name
		users[k].IsFollow = true
	}
	str := "获取关注的人成功"
	followlist := FollowList{
		StatusCode: "0",
		StatusMsg:  &str,
		UserList:   users,
	}
	c.JSON(200, followlist)
}
