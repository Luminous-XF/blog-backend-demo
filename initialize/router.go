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

	publicGroupV1 := router.Group("v1")
	{
		routers.InitBaseRouter(publicGroupV1)
		routers.InitUserPublicRouter(publicGroupV1)
		routers.InitPostPublicRouter(publicGroupV1)
	}

	privateGroupV1 := router.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		routers.InitUserPrivateRouter(privateGroupV1)
		routers.InitPostPrivateRouter(privateGroupV1)
	}

	return router
}
