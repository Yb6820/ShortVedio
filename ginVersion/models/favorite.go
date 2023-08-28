package models

import (
	"DouYin/utils"

	"gorm.io/gorm"
)

type Favorite struct {
	UserId     uint `gorm:"primaryKey"`
	VideoId    uint `gorm:"primaryKey"`
	IsFavorite bool
}

func (table *Favorite) TableName() string {
	return "favorite"
}

// 创建用户点赞某视频的关系
func CreateFavorite(favorite Favorite) *gorm.DB {
	return utils.DB.Create(&favorite)
}

// 更新用户点赞和取消点赞的操作
func UpdateFavorite(favorite Favorite) *gorm.DB {
	favoritesql := Favorite{}
	utils.DB.Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).First(&favoritesql)
	return utils.DB.Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).Model(favorite).Update("is_favorite", !favoritesql.IsFavorite)
}

// 根据用户Id返回它点过赞的作品
func GetFavoriteListByUserId(userid uint) []Favorite {
	data := make([]Favorite, 0)
	utils.DB.Find(&data, "user_id = ? and is_favorite = true", userid)
	return data
}

// 获取用户与作者之间的点赞信息
func IsFavoriteOrNot(userid uint, videoid uint) bool {
	favorite := Favorite{}
	utils.DB.Where("user_id = ? and video_id = ?", userid, videoid).First(&favorite)
	return favorite.IsFavorite
}

// 确认点赞信息是否已经入库
func GetFavoriteById(videoid uint, userid uint) Favorite {
	favorite := Favorite{}
	utils.DB.Where("video_id = ? and user_id = ?", videoid, userid).First(&favorite)
	return favorite
}
