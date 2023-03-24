package router

import (
	"github.com/gin-gonic/gin"
	"talkon_api/user_web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("user")
	{
		UserGroup.GET("list", api.GetUserList)
	}
}
