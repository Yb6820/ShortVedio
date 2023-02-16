package models

import (
	"DouYin/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string
	Password string
	Token    string
	Follow   int32
	Follower int32
	Salt     string
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
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
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
