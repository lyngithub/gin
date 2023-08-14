package console

import (
	"fmt"
	"github.com/robfig/cron"
)

func InitConsole() {
	var c = cron.New()

	err := c.AddFunc("*/1 * * * *", func() {
		//domain.GetOnline()
		fmt.Println("This task runs every minute.")
	})
	if err != nil {
		return
	}

	c.Start()
}
