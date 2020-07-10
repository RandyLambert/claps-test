package common

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err!=nil {
		log.Println(err.Error())
	}

	viper.SetConfigType("json")
	viper.SetConfigName(viper.GetString("MIXIN_CLIENT_CONFIG"))
	//两个配置文件合并
	err = viper.MergeInConfig()
	if err != nil {
		log.Println(err.Error())
	}


	//获取数据库配置信息
	database := viper.GetString("DATABASE_CONFIG")
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(database),&dat); err != nil{
		log.Println(err.Error())
	}
	viper.Set("host",dat["host"])
	viper.Set("port",dat["port"])
	viper.Set("username",dat["username"])
	viper.Set("password",dat["password"])
	viper.Set("database",dat["database"])

	//1.commom.Loger()改成全局
	//2.注释加上
	//3.database配置改成env类型
	//4.Level 配置

}