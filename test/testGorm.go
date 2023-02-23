package main

import (
	"DouYin/models"

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
	//db.Create(&models.Video{AuthorId: 1, CommentCount: 0, FavoriteCount: 0, PlayURL: "https//v26-web.douyinvod.com/1771486aa145ba426cd7f354a89a6227/63f7053a/video/tos/cn/tos-cn-ve-15c001-alinc2/oMh9AlU1A3TcUmz0BgNgfvANQtIwMBwcAtej4A/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1303&bt=1303&cs=0&ds=6&ft=bvTKJbQQqUYqfJEZPo0OW_EklpPiXFE.SOVJEUTgMQCPD-I&mime_type=video_mp4&qs=0&rc=OGZoOzxoOjQzZWQ8OTg3OEBpM2o2cTM6Zmh1aTMzNGkzM0AzNC9hXjNfXmIxXmAvYmFfYSNkNi1rcjQwYW5gLS1kLWFzcw%3D%3D&l=202302231318251A4AE608F5EBBA0208D3&btag=8000", Title: "手机掉了，谁帮我捡一下!", CoverURL: "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/9996399a555446e0a301908215208cda_1677047504~tplv-dy-cropcenter:323:430.jpeg?biz_tag=pcweb_cover&from=3213915784&s=PackSourceEnum_PUBLISH&sc=cover&se=true&sh=323_430&x-expires=1992488400&x-signature=%2BQQyXdkSxOPv7t2HnxJx8dB6U0Y%3D"})
	video := models.Video{}
	db.Where("id = ?", 1).First(&video)
	db.Model(&video).Update("favorite_count", 3)
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
