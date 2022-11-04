package cronJob

import "github.com/robfig/cron/v3"

func InitCron() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * * * * *", func() {
		print("test")
	})
	c.Start()
}
