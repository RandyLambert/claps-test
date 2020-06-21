package common

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	viper.SetConfigType("yml")
	viper.SetConfigName("application")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}