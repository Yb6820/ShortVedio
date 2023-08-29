package main

import (
	"DouYin/router"
	"DouYin/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(viper.GetString("app.port"))
}
