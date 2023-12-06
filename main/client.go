package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	spec := "*/30 * * * * *" // 每隔30s执行一次，cron格式（秒，分，时，天，月，周）
	// 添加一个任务
	c.AddFunc(spec, func() {
		fmt.Println("cron start at : ", time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()

	for true {
		time.Sleep(10 * time.Second)
	}
}
