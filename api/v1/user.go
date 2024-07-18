package v1

import (
	"blog-backend/common/error_code"
	"blog-backend/global"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func CreateTokenByUsernamePassword(ctx *gin.Context) {
	var requestData request.LoginByUsernameAndPasswordRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		global.Logger.Errorf("TraceID:'%s' ErrorCode:%d ErrorInfo:'%s'",
			ctx.GetHeader("Trace-Id"),
			error_code.ParamBindError,
			err.Error(),
		)
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.LoginByUsernameAndPassword(requestData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.Created(responseData, error_code.ErrMsg(code), ctx)
	}
}

func SendVerifyCodeWithEmail(ctx *gin.Context) {
	var requestData request.SendVerifyCodeWithEmailRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		global.Logger.Errorf("TraceID:'%s' ErrorCode:%d ErrorInfo:'%s'",
			ctx.GetHeader("Trace-Id"),
			error_code.ParamBindError,
			err.Error(),
		)
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.SendVerifyCodeWithEmail(requestData, ctx.GetHeader("Trace-Id")); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, responseData, error_code.ErrMsg(code), ctx)
	}
}
