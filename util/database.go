package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
	"time"
)

var DB *gorm.DB

func InitDB() (db *gorm.DB) {
	driverName := "postgres"
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.GetString("DATABASE_HOST"),
		config.GetString("DATABASE_PORT"),
		config.GetString("DATABASE_USERNAME"),
		config.GetString("DATABASE_DATABASE"),
		config.GetString("DATABASE_PASSWORD"))

	DB, err := gorm.Open(driverName, args)

	if err != nil {
		log.Panic("failed to connect database,err :" + err.Error())
		return
	}

	DB.DB().SetConnMaxLifetime(time.Hour)
	DB.DB().SetMaxOpenConns(1024)
	DB.DB().SetMaxIdleConns(32)

	DB.SingularTable(true)

	return DB
}

/*
获取数据库句柄
*/
func GetDB() *gorm.DB {
	return DB
}
