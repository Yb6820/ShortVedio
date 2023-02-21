package models

import (
	"DouYin/utils"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId uint
	UserId  uint
	Content string
}

func (table *Comment) TableName() string {
	return "comment"
}

// 获取视频底下的所有评论
func GetCommentByVideoId(videoid uint) []Comment {
	data := make([]Comment, 10)
	utils.DB.Find(&data, "video_id = ?", videoid)
	return data
}
