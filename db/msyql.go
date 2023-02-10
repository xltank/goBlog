package db

import (
	"goBlog/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitMysql() {
	db, err := gorm.Open(mysql.Open(config.Conf.Mysql.Pass+config.Conf.Mysql.Uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic("connect mysql error:" + err.Error())
	}
	log.Println("MySql connected at", config.Conf.Mysql.Uri)
	Db = db
}
