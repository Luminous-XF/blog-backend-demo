package v1

import (
	"blog-backend/common/error_code"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginForm request.LoginByUsernameAndPasswordRequest
	if err := ctx.ShouldBindBodyWithJSON(&loginForm); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if code := service.LoginByUsernameAndPassword(loginForm); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.SuccessWithMessage(code, error_code.ErrMsg(code), ctx)
	}
}
