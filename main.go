package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xxblog/base/conf"
	log "xxblog/base/logger"
	"xxblog/model"
	"xxblog/repositories"
	"xxblog/routers"
)

func main() {
	conf.Init()
	r := gin.Default()
	repositories.InitDBConnection(model.Models...)
	routers.Init(r)
	log.Info("server start")
	err := r.Run(":8089") // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}
}
