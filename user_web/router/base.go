package router

import (
	"github.com/gin-gonic/gin"
	"talkon_api/user_web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	// 登录接口 不需要token验证
	loginGroup := Router.Group("login")
	{
		loginGroup.POST("email", api.EmailLogin)
	}

	// 验证码接口 不需要token验证
	UserGroup := Router.Group("captcha")
	{
		UserGroup.GET("img", api.GetPicCaptcha)
	}
}
