package main

import (
	"fmt"
	"go.uber.org/zap"
	"talkon_api/user_web/initialize"
)

func main() {
	// 1.初始化 Routers
	eg := initialize.Routers()
	port := "8098"
	// 2.初始化全局 Logger
	initialize.InitLogger()
	// 启动gin服务
	zap.S().Infof("启动服务器，端口：%s", port)
	if err := eg.Run(fmt.Sprintf(":%s", port)); err != nil {
		zap.S().Panic("服务器启动失败：", err.Error())
	}
}
