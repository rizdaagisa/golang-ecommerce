package scheduler

import (
	"fmt"
	cron "github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Run Every 1 minute on gin")
	})

	c.Start()

}
