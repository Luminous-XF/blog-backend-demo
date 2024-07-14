package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    ErrorCode   `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
}

func Unauthorized(code ErrorCode, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, Response{
		Code:    code,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func PermissionDenied(code ErrorCode, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, Response{
		Code:    code,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, Response{
		Code:    ERROR,
		Msg:     "404 not found",
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func Forbidden(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, Response{
		Msg:     "",
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func DeleteSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, Response{
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func CommonFailed(code ErrorCode, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func CommonSuccess(code ErrorCode, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func SuccessWithMessage(code ErrorCode, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func Created(data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, Response{
		Code:    SUCCESS,
		Data:    data,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func Accepted(fileUploadStatus interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, Response{
		Code:    http.StatusAccepted,
		Data:    fileUploadStatus,
		Msg:     "File upload accepted",
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}

func Result(code ErrorCode, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceId: ctx.GetHeader("Trace-Id"),
	})
}
