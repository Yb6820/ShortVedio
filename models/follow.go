package models

import (
	"DouYin/utils"

	"gorm.io/gorm"
)

type Follow struct {
	UserId   uint `gorm:"primaryKey"`
	AuthorId uint `gorm:"primaryKey"`
	IsFollow bool
}

func (table *Follow) TableName() string {
	return "follow"
}

// 创建用户与用户之间的关注与被关注的关系
func CreateFollow(follow Follow) *gorm.DB {
	return utils.DB.Create(&follow)
}

// 更新用户关注和取消关注的操作
func UpdateFollow(follow Follow) *gorm.DB {
	followsql := Follow{}
	utils.DB.Where("user_id = ? and author_id = ?", follow.UserId, follow.AuthorId).First(&followsql)
	return utils.DB.Where("user_id = ? and author_id = ?", follow.UserId, follow.AuthorId).Updates(Favorite{IsFavorite: !followsql.IsFollow})
}

func IsInFollowTable(userid uint, authorid uint) bool {
	res := Follow{}
	utils.DB.Where("user_id = ? and author_id = ?", userid, authorid).First(&res)
	if res.UserId == 0 {
		return false
	} else {
		return true
	}
}

// 获取用户与作者之间的关注信息
func IsFollowOrNot(userid uint, authorid uint) bool {
	follow := Follow{}
	utils.DB.Where("user_id = ? and author_id = ?", userid, authorid).First(follow)
	return follow.IsFollow
}

// 获取某个用户关注的所有作者
func GetFollows(userid uint) []Follow {
	data := make([]Follow, 10)
	utils.DB.Find(&data, "user_id = ? and is_follow = true", userid)
	return data
}

// 获取某个用户的所有关注者
func GetFollowers(authorid uint) []Follow {
	data := make([]Follow, 10)
	utils.DB.Find(&data, "author_id = ? and is_follow = true", authorid)
	return data
}
