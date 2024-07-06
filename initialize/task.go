package initialize

import (
	"fmt"
	"xx/console"
)

func init() {
	c := console.StartCronJob()
	fmt.Println(c.ErrorLog)
	//defer c.Stop()
}
