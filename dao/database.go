package dao

import (
	"claps-test/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := viper.GetString("host")
	port := viper.GetString("port")
	database := viper.GetString("database")
	username := viper.GetString("username")
	password := viper.GetString("password")
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
		common.Logger().Panic("failed to connect database,err :" + err.Error())
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
