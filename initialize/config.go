package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"xx/global"
)

func InitConfig() {

	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err.Error())
	}

	fmt.Println(global.ServerConfig)

}
