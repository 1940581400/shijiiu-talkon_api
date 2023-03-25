package initialize

import (
	"log"
	"strconv"
	"talkon_api/user_web/config"
	"talkon_api/user_web/global"

	"github.com/spf13/viper"
)

func GetEnvInfo(key string) string {
	viper.AutomaticEnv()
	return viper.GetString(key)
}

func InitConfig() {
	configName := "user_web/config_pro.yaml"
	info := GetEnvInfo("TALKON_DEV")
	if flag, _ := strconv.ParseBool(info); flag {
		configName = "user_web/config_dev.yaml"
	}
	logger := log.Default()
	logger.Printf("[配置文件] 读取中 %s", configName)
	v := viper.New()
	v.SetConfigFile(configName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	global.ServerConfig = &serverConfig
	logger.Printf("[配置文件] 读取完成")
}
