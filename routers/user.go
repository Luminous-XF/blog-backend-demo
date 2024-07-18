package routers

import (
	v1 "blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserPublicRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.POST("/uuid", v1.CreateTokenByUsernamePassword)
		UserRouter.POST("/verify-code/using-email", v1.SendVerifyCodeWithEmail)
	}
}

func InitUserPrivateRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.PUT("/uuid", v1.CreateTokenByUsernamePassword)
	}
}
