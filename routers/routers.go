package routers

import (
	"chuxin0816/SE/controller"
	"chuxin0816/SE/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRouter := r.Group("/categories")
	{
		categoryController := controller.NewCategoryController()
		categoryRouter.POST("", categoryController.Create)
		categoryRouter.PUT("/:id", categoryController.Update)
		categoryRouter.GET("/:id", categoryController.Show)
		categoryRouter.DELETE("/:id", categoryController.Delete)
	}

	postRouter := r.Group("/posts")
	{
		postController := controller.NewPostController()
		postRouter.Use(middleware.AuthMiddleware())
		postRouter.POST("", postController.Create)
		postRouter.PUT("/:id", postController.Update)
		postRouter.GET("/:id", postController.Show)
		postRouter.DELETE("/:id", postController.Delete)
		postRouter.GET("/page/list",postController.PageList)
	}
	return r
}
