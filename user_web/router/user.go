package router

import (
	"github.com/gin-gonic/gin"
	"talkon_api/user_web/api"
	"talkon_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {

	// 用户接口需要token验证
	UserGroup := Router.Group("user").Use(middlewares.JwtTokenAuth())
	{
		UserGroup.POST("list", api.GetUserList)
	}
}
