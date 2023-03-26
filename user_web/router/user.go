package router

import (
	"github.com/gin-gonic/gin"
	"talkon_api/user_web/api"
	"talkon_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// 登录接口不需要token验证
	loginGroup := Router.Group("login")
	{
		loginGroup.POST("email", api.EmailLogin)
	}
	// 用户接口需要token验证
	UserGroup := Router.Group("user").Use(middlewares.JwtTokenAuth())
	{
		UserGroup.POST("list", api.GetUserList)
	}
}
