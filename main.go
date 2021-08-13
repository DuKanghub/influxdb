package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"time"
	"w2influx/service"
)


func main() {
	c := cron.New()
	// 定时任务这里的时间是秒 分 时 天 月 周
	c.AddFunc("*/10 * * * * *", service.DbSize2Influx)
	c.Start()
	var cmd string
	for {
		time.Sleep(2*time.Second)
		cmd = os.Getenv("cron")
		fmt.Println(cmd)
		if  cmd == "stop" {
			c.Stop()
			break
		}
	}
	os.Exit(0)
}