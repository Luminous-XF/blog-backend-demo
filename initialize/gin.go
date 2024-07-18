package initialize

import (
	"blog-backend/middleware"
	"blog-backend/routers"
	"github.com/gin-gonic/gin"
)

func initGin() (engine *gin.Engine) {
	engine = gin.Default()
	initMiddleware(engine)
	initRouters(engine)
	return engine
}

func initMiddleware(engine *gin.Engine) {
	engine.Use(middleware.Cors())
	engine.Use(middleware.AddTracID())
}

func initRouters(engine *gin.Engine) {
	publicGroupV1 := engine.Group("v1")
	{
		routers.InitBaseRouter(publicGroupV1)
		routers.InitUserPublicRouter(publicGroupV1)
		routers.InitPostPublicRouter(publicGroupV1)
	}

	privateGroupV1 := engine.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		routers.InitUserPrivateRouter(privateGroupV1)
		routers.InitPostPrivateRouter(privateGroupV1)
	}
}
