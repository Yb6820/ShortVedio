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

// 创建评论信息
func CreateComment(comment Comment) *gorm.DB {
	return utils.DB.Create(&comment)
}

// 获取视频底下的所有评论
func GetCommentByVideoId(videoid uint) []Comment {
	data := make([]Comment, 10)
	utils.DB.Find(&data, "video_id = ?", videoid)
	return data
}

// 根据commentid获取Comment信息
func GetCommentByCommentId(commentid uint) Comment {
	comment := Comment{}
	utils.DB.Where("comment_id = ?", commentid).First(comment)
	return comment
}

// 逻辑删除评论数据
func DeleteCommentById(comment Comment) *gorm.DB {
	return utils.DB.Delete(comment)
}
