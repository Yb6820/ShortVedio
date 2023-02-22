package models

import (
	"DouYin/utils"
	"fmt"

	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     int  //对应的类型 1好友  2群  3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func CreateContact(contact Contact) *gorm.DB {
	return utils.DB.Create(&contact)
}
func SearchFriends(userId uint) []uint {
	contacts := make([]Contact, 0)
	objIds := make([]uint, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(v)
		objIds = append(objIds, v.TargetId)
	}
	return objIds
}
