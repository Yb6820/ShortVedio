package models

import (
	"DouYin/utils"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      uint   // 视频作者信息
	CommentCount  int64  // 视频的评论总数
	CoverURL      string // 视频封面地址
	FavoriteCount int64  // 视频的点赞总数
	PlayURL       string // 视频播放地址
	Title         string // 视频标题
}

func (table *Video) TableName() string {
	return "video"
}

// 创建视频在数据库中的存储样式
func CreateVideo(video Video) *gorm.DB {
	return utils.DB.Create(&video)
}

// 根据作者Id获取所有视频信息
func GetVideosByAuthorId(authorId uint) []Video {
	data := make([]Video, 10)
	utils.DB.Find(&data, "author_id = ?", authorId)
	return data
}

// 通过VideoId获取视频信息
func GetVideoById(videoid uint) Video {
	video := Video{}
	utils.DB.Where("id = ?", videoid).First(&video)
	return video
}

// 获取在某时间节点之前的所有视频
func GetVideosBeforeTime(time string) []Video {
	videos := make([]Video, 0)
	utils.DB.Find(&videos, "created_at > ?", time)
	return videos
}

// 实现点赞操作计数
func UpdateFavoriteCount(video Video, flag bool) *gorm.DB {
	if flag {
		video.FavoriteCount += 1
	} else {
		video.FavoriteCount -= 1
	}
	return utils.DB.Model(&video).Update("favorite_count", video.FavoriteCount)
}

// 实现评论操作计数
func UpdateCommentCount(video Video, flag bool) *gorm.DB {
	if flag {
		video.CommentCount += 1
	} else {
		video.CommentCount -= 1
	}
	return utils.DB.Model(&video).Update("comment_count", video.CommentCount)
}
