package models

import (
	"DouYin/utils"

	"gorm.io/gorm"
)

type Follow struct {
	UserId   uint
	AuthorId uint
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

// 获取用户与作者之间的关注信息
func IsFollowOrNot(userid uint, authorid uint) bool {
	follow := Follow{}
	utils.DB.Where("user_id = ? and author_id = ?", userid, authorid).First(follow)
	return follow.IsFollow
}
