package main

import (
	"github.com/robfig/cron/v3"
	"proxy/client"
	"proxy/cons"
	"proxy/tools"
	"time"
)

func main() {

	tools.InitViper()
	cons.InitConst()
	tools.InitLogger()

	c := cron.New(cron.WithSeconds())
	spec := "0 */5 * * * ?" // 每隔5m执行一次，cron格式（秒，分，时，天，月，周）
	// 添加一个任务
	c.AddFunc(spec, func() {
		tools.Logger.Info("cron start at : ", time.Now().Format("2006-01-02 15:04:05"))
		client.CheckJob()

	})
	c.Start()

	for true {
		time.Sleep(10 * time.Second)
	}
}
