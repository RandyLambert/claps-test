package common

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	//workDir, err := os.Getwd()
	//if err != nil {
	//	panic(err.Error())
	//}
	//viper.SetConfigType("yml")
	//viper.SetConfigName("application")
	//viper.AddConfigPath(workDir + "/config")
	//err = viper.ReadInConfig()
	//if err != nil {
	//	panic(err.Error())
	//}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err!=nil {
		log.Println(err.Error())
	}

	viper.SetConfigType("json")
	viper.SetConfigName(viper.GetString("MIXIN_CLIENT_CONFIG"))
	err = viper.MergeInConfig()
	if err != nil {
		log.Println(err.Error())
	}
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

}