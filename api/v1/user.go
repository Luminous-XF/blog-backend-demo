package v1

import (
	"blog-backend/common/error_code"
	"blog-backend/global"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var formData request.LoginByUsernameAndPasswordRequest
	if err := ctx.ShouldBindBodyWithJSON(&formData); err != nil {
		global.Logger.Errorf("TraceID:'%s' ErrorCode:%d ErrorInfo:'%s'",
			ctx.GetHeader("Trace-Id"),
			error_code.ParamBindError,
			err.Error(),
		)
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if code := service.LoginByUsernameAndPassword(formData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.SuccessWithMessage(code, error_code.ErrMsg(code), ctx)
	}
}
