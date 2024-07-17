package routers

import (
	v1 "blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.POST("/token/username-password", v1.CreateTokenByUsernamePassword)
}
