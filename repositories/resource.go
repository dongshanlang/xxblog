package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xxblog/base/conf"
	log "xxblog/base/logger"
)

var DB *gorm.DB

func InitDBConnection(models ...interface{}) {
	var err error
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	if DB, err = gorm.Open("mysql", conf.Conf.Mysql.DSN); err != nil {
		log.Errorf("opens database failed: %s", err.Error())
		panic(err)
		return
	}

	DB.LogMode(true)
	DB.SingularTable(true) // 禁用表名负数
	DB.DB().SetMaxIdleConns(conf.Conf.Mysql.MaxIdleConnection)
	DB.DB().SetMaxOpenConns(conf.Conf.Mysql.MaxConnection)

	if conf.Conf.Init.DB {
		if err = DB.AutoMigrate(models...).Error; nil != err {
			log.Errorf("auto migrate tables failed: %s", err.Error())
			panic(err)
		}
	}
	log.Info("database connect ok")
	return
}
