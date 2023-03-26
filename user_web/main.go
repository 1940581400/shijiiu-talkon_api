package main

import (
	"talkon_api/user_web/initialize"
)

func main() {
	// 1.初始化配置，注意 此操作一定要在第一，不然后面初始化读不到配置
	initialize.InitConfig()
	// 2.初始化日志，注意 此操作一定要在第二，不然初始化文件当中的 日志 无法打印
	initialize.InitLogger()
	// 3.注册翻译器，注意 此操作一定要在第三，不然后面的自定义验证器取不到全局翻译器
	initialize.InitTrans("zh")

	// 下面顺序可随意
	// 4.初始化自定义验证器
	initialize.InitValidators()

	// 启动gin 一定要在最后，因为它会一直阻塞
	initialize.StartServer()

}
