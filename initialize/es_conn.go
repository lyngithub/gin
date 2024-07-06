package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xx/global"
	"xx/utils/es"
)

func init() {
	err := es.InitClient(
		es.DefaultClient,
		[]string{fmt.Sprintf("%s:%d", global.ServerConfig.EsInfo.Addr, global.ServerConfig.EsInfo.Port)},
		"",
		"")
	if err != nil {
		log.Println(err)
		return
	}
}
