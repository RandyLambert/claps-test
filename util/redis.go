package util

import (
	"github.com/go-redis/redis"
	config "github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("REDIS_ADDR"),
		Password: config.GetString("REDIS_PASSWORD"), // no password set
		DB:       config.GetInt("REDIS_DB"),  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}