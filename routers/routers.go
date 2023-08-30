package routers

import (
	"chuxin0816/SE/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)
	return r
}
