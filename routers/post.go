package routers

import (
	v1 "blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPostRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("posts")
	{
		UserRouter.POST("", v1.GetPostList)
	}
}
