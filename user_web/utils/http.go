package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func GrpcErrToHttpErr(err error, ctx *gin.Context) {
	if err == nil {
		return
	}
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": e.Message(),
			})
		case codes.Unimplemented:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无权操作",
			})
		case codes.Internal:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "系统内部异常",
			})
		case codes.Unavailable:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务不可用",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "其他异常",
			})
		}
	}
}
