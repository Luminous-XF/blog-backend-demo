package routers

import (
	v1 "blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.POST("/login", v1.Login)
	}
}
