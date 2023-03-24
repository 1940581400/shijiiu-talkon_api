package main

import (
	"time"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"./user_web.log",
		"stderr",
	}
	return config.Build()
}

func main() {
	//logger, _ := zap.NewProduction() //生产环境
	// logger, _ := zap.NewDevelopment() // 测试环境
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "shijiu.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	logger.Info("failed to fetch URL",
		zap.String("url", url),
	)
}
