package v1

import (
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginForm request.LoginByUsernamePassword
	if err := ctx.ShouldBind(&loginForm); err != nil {
		response.CommonFailed(response.ParamBindError, response.ErrMsg(response.ParamBindError), ctx)
		return
	}

	if code := service.LoginByUsernameAndPassword(loginForm); response.IsSuccess(code) {
		response.CommonFailed(code, response.ErrMsg(code), ctx)
	} else {
		response.SuccessWithMessage(code, response.ErrMsg(code), ctx)
	}
}
