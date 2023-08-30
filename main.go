package main

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/routers"
)

func main() {
	common.InitDB()
	common.DB.AutoMigrate(&models.User{})

	r := routers.SetupRouter()
	panic(r.Run())
}
