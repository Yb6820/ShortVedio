package main

import (
	"DouYin/router"
	"DouYin/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run()
}
