package main

import (
	"DouYin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	db, err := gorm.Open(mysql.Open("root:jun181234@tcp(121.37.246.78:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//创建表，没有则新创
	//db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.UserBasic{}, &models.Comment{}, &models.Contact{}, &models.Favorite{}, &models.Follow{}, &models.Video{})
	//创建视频信息
	//db.Create(&models.Video{AuthorId: 1, CommentCount: 0, FavoriteCount: 0, PlayURL: "http://v26-web.douyinvod.com/c69a4069db38e55e0050ee11af56aa46/63f78d8a/video/tos/cn/tos-cn-ve-15/oUOICUgxgABbA9BBeftEMxzBJB3jAhAHPHjbQ5/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1407&bt=1407&cs=0&ds=3&ft=bvTKJbQQqUYqfJEZPo0OW_EklpPiXwOySOVJEUTgMQCPD-I&mime_type=video_mp4&qs=0&rc=OjNoOjY2OzM0ZjU0PDg0OUBpMzRmazQ6ZmdyaDMzNGkzM0BgMmMzL18wNV8xLjA1NmNjYSMtYjFycjRvLnBgLS1kLS9zcw%3D%3D&l=20230223230003FCF7B4F36FC91E203AAB&btag=8000", Title: "高三牲!", CoverURL: "https://p6-pc-sign.douyinpic.com/tos-cn-p-0015/326f96ef143e46aebd8f81a6e59dca0e_1673101545~tplv-dy-cropcenter:323:430.jpeg?biz_tag=pcweb_cover&from=3213915784&s=PackSourceEnum_PUBLISH&sc=cover&se=true&sh=323_430&x-expires=1992524400&x-signature=jRQiVWNEH%2BZxfapirtaok84%2FmwY%3D"})
	//db.Create(&models.Video{AuthorId: 2, CommentCount: 0, FavoriteCount: 0, PlayURL: "http://v26-web.douyinvod.com/f99a8843ef15fb3f66832a9e0481ec2f/63f796ad/video/tos/cn/tos-cn-ve-15c001-alinc2/oIfBsePl5Nhy40l3p7aIAC4BeaNhREA2AH3o1H/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1718&bt=1718&cs=0&ds=4&ft=bvTKJbQQqUYqfJEZPo0OW_EklpPiXu_HSOVJEUTgMQCPD-I&mime_type=video_mp4&qs=0&rc=aTM2ZzU2Nzs0Zzg1MzU2NUBpMztndGU6Zjl1aTMzNGkzM0BhNDQ2YDEtNjExMS9eNF5jYSM2cGRfcjRfb2xgLS1kLS9zcw%3D%3D&l=20230223233852EDEDD2EFAC3E9C23282B&btag=10000", Title: "自发光而不是被照亮", CoverURL: "https://p6-pc-sign.douyinpic.com/tos-cn-p-0015/bc939a80a43746f38715fadfe5881429_1676787109~tplv-dy-cropcenter:323:430.jpeg?biz_tag=pcweb_cover&from=3213915784&s=PackSourceEnum_PUBLISH&sc=cover&se=true&sh=323_430&x-expires=1992524400&x-signature=z%2B8IpU2IfcRmrlWTPEmHXaF2z50%3D"})
	//db.Create(&models.Video{AuthorId: 1, CommentCount: 0, FavoriteCount: 0, PlayURL: "http://v26-web.douyinvod.com/b23ccc10e87f0bf318a73081e1a530ab/63f796a2/video/tos/cn/tos-cn-ve-15c001-alinc2/o0NCc2caTx8wIghzc07fPAFzBkBA3Ue6jDDAQA/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=2189&bt=2189&cs=0&ds=3&ft=bvTKJbQQqUYqfJEZPo0OW_EklpPiXu_HSOVJEUTgMQCPD-I&mime_type=video_mp4&qs=0&rc=OWhkZTw7ZzU5aGY3PGYzZkBpM3BteTY6ZjQ6aDMzNGkzM0AtMjUzYDAvNl8xYjFeX14uYSMtazM0cjQwMHFgLS1kLTBzcw%3D%3D&l=20230223233852EDEDD2EFAC3E9C23282B&btag=8000", Title: "满脑子都是Hey brother", CoverURL: "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/64a2c5eaa8a14850a9d548b25f1169c9_1673163205~tplv-dy-cropcenter:323:430.jpeg?biz_tag=pcweb_cover&from=3213915784&s=PackSourceEnum_PUBLISH&sc=cover&se=true&sh=323_430&x-expires=1992524400&x-signature=jE%2FAtXNQgDXc71ohw6rYhtUlGs8%3D"})

	//创建新表
	//db.AutoMigrate(&Student{})

	//插入数据
	/* stu := Student{
		Name: "张三",
		Age:  18,
		Sex:  true,
	}
	//db.Create(&stu)
	students := make([]Student, 2)
	students[0] = stu
	stu.Name = "李四"
	students[1] = stu
	db.Create(&students) */

	//查询数据
	/* stu := make([]Student, 0)
	db.Find(&stu, "age = ? ", 18)
	fmt.Println(stu) */

	//更新数据
	//stu := Student{}
	//db.Model(&stu).Where("name = ?", "张三").Update("age", 20)

	//删除数据
	stu := Student{}
	db.Where("name = ?", "张三").Delete(&stu)

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
