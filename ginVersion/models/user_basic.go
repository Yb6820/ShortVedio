package models

import (
	"DouYin/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name            string //姓名
	Password        string //密码
	Token           string //用户登录token
	Avatar          string // 用户头像
	BackgroundImage string // 用户个人页顶部大图
	Follow          int64  // 关注总数
	Follower        int64  // 粉丝总数
	FavoriteCount   int64  // 喜欢数
	WorkCount       int64  // 作品数
	TotalFavorited  string // 获赞数量
	Signature       string // 个人简介
}

// 存放在ctx中的用户信息的模型
type UserInfo struct {
	Username string `json:"username"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 名字唯一确定一个人
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}
func CreateUser(user UserBasic) (err error) {
	err = utils.DB.Create(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func FindUserByID(Id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", Id).First(&user)
	return user
}

// 登录校验
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and password = ?", name, password).First(&user)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("token", temp)
	return user
}
