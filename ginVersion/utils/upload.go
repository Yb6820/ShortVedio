package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"mime/multipart"
)

var (
	accessKey = viper.GetString("qiNiu.accessKey")
	secretKey = viper.GetString("qiNiu.secretKey")
	bucket    = viper.GetString("qiNiu.bucket")
	zone      = viper.GetInt("qiNiu.zone")
	imgUrl    = viper.GetString("qiNiu.imgUrl")
)

func UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          selectZone(zone),
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	//上传文件
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	//打开文件流
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	err = formUploader.Put(ctx, &ret, upToken, "douyin/"+fileHeader.Filename, file, fileHeader.Size, &putExtra)
	if err != nil {
		return "", err
	}
	url = imgUrl + ret.Key
	return url, nil
}
func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuadongZheJiang2
	case 3:
		return &storage.ZoneHuabei
	case 4:
		return &storage.ZoneHuanan
	case 5:
		return &storage.ZoneBeimei
	case 6:
		return &storage.ZoneXinjiapo
	case 7:
		return &storage.ZoneShouEr1
	default:
		return &storage.ZoneHuanan
	}
}
