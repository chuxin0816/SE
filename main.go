package main

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/routers"
)

func main() {
	common.InitDB()

	r := routers.SetupRouter()
	panic(r.Run())
}
