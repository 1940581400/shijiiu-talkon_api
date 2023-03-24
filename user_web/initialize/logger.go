package initialize

import "go.uber.org/zap"

func InitLogger() {
	// 生产环境logger，级别为info，debug日志不会打印
	//logger, err := zap.NewProduction()
	// 测试环境logger，级别为debug
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}
