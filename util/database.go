package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := viper.GetString("DATABASE_HOST")
	port := viper.GetString("DATABASE_PORT")
	database := viper.GetString("DATABASE_DATABASE")
	username := viper.GetString("DATABASE_USERNAME")
	password := viper.GetString("DATABASE_PASSWORD")
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)

	if err != nil {
		log.Panic("failed to connect database,err :" + err.Error())
	}

	db.SingularTable(true)

	DB = db
	return db
}

/*
获取数据库句柄
*/
func GetDB() *gorm.DB {
	return DB
}
