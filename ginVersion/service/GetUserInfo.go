package service

import (
	"DouYin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Userbasic models.UserBasic

type UserInfo struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	User       *User   `json:"user"`        // 用户信息
}

/*
// User
type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
*/

// User
type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

// 根据用户ID返回json格式的User数据
func GetUserInfoById(userid uint, authorid uint) User {
	usersql := models.FindUserByID(userid)
	user := User{
		ID:            int64(usersql.ID),
		FollowCount:   usersql.Follow,
		FollowerCount: usersql.Follower,
		IsFollow:      models.IsFollowOrNot(userid, authorid), //待后续修改
		Name:          usersql.Name,
	}
	return user
}

func GetUserInfo(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	user := GetUserInfoById(uint(user_id), uint(user_id))
	str := "获取用户信息成功"
	userinfo := UserInfo{
		StatusCode: 0,
		StatusMsg:  &str,
		User:       &user,
	}
	c.JSON(200, userinfo)
}
