package main

import (
	"fmt"
	"xx/global"
	"xx/initialize"
)

func main() {
	router := initialize.Routers()

	if err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		panic("启动失败")
	} else {
		fmt.Println("成功")
	}

}
