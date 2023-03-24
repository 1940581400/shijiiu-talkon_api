package initialize

import (
	"github.com/gin-gonic/gin"
	"talkon_api/user_web/router"
)

func Routers() *gin.Engine {
	eg := gin.Default()
	AipGroup := eg.Group("u/v1")
	router.InitUserRouter(AipGroup)
	return eg
}
