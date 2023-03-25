package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"talkon_api/user_web/global"
)

// InitServer 启动gin服务
func InitServer() {
	eg := Routers()
	config := global.ServerConfig.UserWebInfo
	port := "8098"
	zap.S().Infof("[服务器] 启动，地址：%s 端口：%s", config.Host, config.Port)
	if err := eg.Run(fmt.Sprintf(":%s", port)); err != nil {
		zap.S().Panic("[服务器] 启动失败：", err.Error())
	}
}
