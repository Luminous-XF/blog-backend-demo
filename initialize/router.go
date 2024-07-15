package initialize

import (
	"blog-backend/middleware"
	"blog-backend/routers"
	"github.com/gin-gonic/gin"
)

func initRouters() (router *gin.Engine) {
	router = gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.AddTracID())

	privateGroupV1 := router.Group("v1")
	{
		routers.InitUserRouter(privateGroupV1)
		routers.InitPostRouter(privateGroupV1)
	}

	return router
}
