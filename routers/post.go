package routers

import (
	v1 "blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPostPublicRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("posts")
	{
		UserRouter.POST("", v1.GetPostList)
		UserRouter.POST("/uuid", v1.GetPostInfoByUUID)
	}
}

func InitPostPrivateRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("posts")
	{
		UserRouter.PUT("/uuid", v1.GetPostInfoByUUID)
	}
}
