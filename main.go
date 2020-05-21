package main

import (
	"github.com/gin-gonic/gin"
	"xxblog/conf"
	log "xxblog/logger"
	"xxblog/routers"
)

func main() {
	r := gin.Default()
	initLog()
	routers.Init(r)
	log.Info("server start")
	err := r.Run(":8089") // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}
}
func initLog() {
	c := log.New()
	c.SetDivision("time")    // 设置归档方式，"time"时间归档 "size" 文件大小归档，文件大小等可以在配置文件配置
	c.SetTimeUnit(log.Day)   // 时间归档 可以设置切割单位
	c.SetEncoding("console") // 输出格式 "json" 或者 "console"
	if conf.Conf.Debug {
		c.Debug()
	}
	if !conf.Conf.Log.Stdout {
		c.CloseConsoleDisplay()
	}
	c.SetInfoFile(conf.Conf.Log.Info)
	c.SetErrorFile(conf.Conf.Log.Error)
	c.InitLogger()
}
