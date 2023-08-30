package main

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/routers"
	"os"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()
	common.DB.AutoMigrate(&models.User{})
	r := routers.SetupRouter()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
