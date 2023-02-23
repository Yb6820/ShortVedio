package main

import (
	"DouYin/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:181234@tcp(121.37.246.78:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//创建表，没有则新创
	//db.AutoMigrate(&models.UserBasic{})

	//生成message表
	//db.AutoMigrate(&models.Message{})

	//生成contact表
	//db.AutoMigrate(&models.Comment{})

	//创建视频信息
	//db.Create(&models.Video{AuthorId: 1, CommentCount: 0, FavoriteCount: 0, PlayURL: "https://v26-web.douyinvod.com/05dba22339fdd59db87141b9791504a5/63f74346/video/tos/cn/tos-cn-ve-15c001-alinc2/ow0fOAFsjrAMyI35By8DgknaAvA9tfJRhsbwCn/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1528&bt=1528&cs=0&ds=4&ft=bvTKJbQQqUYqfJEZPo0OW_EklpPiXOSmSOVJEUTgMQCPD-I&mime_type=video_mp4&qs=0&rc=ZmUzNjUzNWloaDxpNGQ6O0BpajVxbWY6Zjk6aDMzNGkzM0AtXmEzX2M1X2ExMS9jY2EvYSMybDZpcjRnZnNgLS1kLS9zcw%3D%3D&l=20230223174311216C4B1FC2517904CF02&btag=8000", Title: "手机掉了，谁帮我捡一下!", CoverURL: "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/73264c8c1c43434aa9c7b73050141c1f_1673428090~tplv-dy-cropcenter:323:430.jpeg?biz_tag=pcweb_cover&from=3213915784&s=PackSourceEnum_PUBLISH&sc=cover&se=true&sh=323_430&x-expires=1992502800&x-signature=gOa3azeTdklbR5BoODoQcx9Ret0%3D"})

	test := models.Favorite{}
	db.Where("user_id = ? and video_id = ?", 1, 2).First(&test)
	fmt.Println(test)

	/* video := models.Video{}
	db.Where("id = ?", 1).First(&video)
	db.Model(&video).Update("favorite_count", 3) */
	//生成group_basic表
	//db.AutoMigrate(&models.Follow{})
	/* x := "2023-02-16 12:48:42.800"
	t, _ := time.Parse("2006-01-02 15:04:05", x)
	fmt.Println(t.Unix())
	user := make([]models.UserBasic, 10)
	tm := time.Now()
	tm.Format("2006-01-02 15:04:05")
	db.Find(&user, "created_at < ?", tm)
	fmt.Println(user) */
	/* // Create
	user := &models.UserBasic{}
	user.Name = "张三"
	user.LoginTime = time.Now()
	user.LogOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	db.Create(user)

	// 读取数据
	fmt.Println(db.First(user, 1)) // find product with integer primary key
	//db.First(user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(user).Update("PassWord", 1234) */
	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
}
