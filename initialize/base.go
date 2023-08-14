package initialize

import "xx/console"

func InitBase() {
	InitConfig()
	InitMysql()
	InitLogger()
	InitRedis()
	console.InitConsole()
	InitListenChan()
}
