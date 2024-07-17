package v1

import (
	"blog-backend/common/error_code"
	"blog-backend/global"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func GetPostList(ctx *gin.Context) {
	var formData request.PageInfoRequest
	if err := ctx.ShouldBindBodyWithJSON(&formData); err != nil {
		global.Logger.Errorf("TraceID:'%s' ErrorCode:%d ErrorInfo:'%s'",
			ctx.GetHeader("Trace-Id"),
			error_code.ParamBindError,
			err.Error(),
		)
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if postList, code := service.GetPostList(formData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, postList, error_code.ErrMsg(code), ctx)
	}
}

func GetPostInfoByUUID(ctx *gin.Context) {
	var formData request.GetByUUIDRequest
	if err := ctx.ShouldBindBodyWithJSON(&formData); err != nil {
		global.Logger.Errorf("TraceID:'%s' ErrorCode:%d ErrorInfo:'%s'",
			ctx.GetHeader("Trace-Id"),
			error_code.ParamBindError,
			err.Error(),
		)
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if post, code := service.GetPostByUUID(formData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, post, error_code.ErrMsg(code), ctx)
	}
}
