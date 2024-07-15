package v1

import (
	"blog-backend/common/error_code"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/service"
	"github.com/gin-gonic/gin"
)

func GetPostList(ctx *gin.Context) {
	var pageInfoForm request.PageInfoRequest
	if err := ctx.ShouldBindBodyWithJSON(&pageInfoForm); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if postList, code := service.GetPostList(pageInfoForm); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, postList, error_code.ErrMsg(code), ctx)
	}
}
