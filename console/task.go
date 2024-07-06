package console

import (
	"fmt"
	"github.com/robfig/cron"
)

func StartCronJob() *cron.Cron {
	var c = cron.New()

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("This task runs every minute.")
	})

	c.Start()
	return c
}
