package main

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func main() {
	putPolicy := storage.PutPolicy{
		Scope: "blogyoubet",
	}
	mac := qbox.NewMac("sHKf9jw62Zh6xqGhpmQsPVzlBAjZfs8ryI3lWrHY", "yf5o_qtrmSYwOv0IUGAbBzLZsd1XUL-KJf6AqtFv")
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, "DouYin/test.jpg", "D:\\桌面\\图片资源\\Dog.jpg", &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
