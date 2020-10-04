package util

import (
	"fmt"
	"github.com/gin-contrib/cache/persistence"
	"time"
)

// 声明一个全局的rdb变量
//var Rdb *persistence.RedisStore
var Rdb *persistence.InMemoryStore

//jwt的过期时间
const TokenExpireDuration = time.Hour * 2

// 初始化连接
func InitClient() (err error) {
	Rdb = persistence.NewInMemoryStore(TokenExpireDuration)
	/*
		Rdb = persistence.NewRedisCache(viper.GetString("REDIS_ADDR"),
			viper.GetString("REDIS_PASSWORD"),TokenExpireDuration)
	*/
	if Rdb == nil {
		fmt.Println("Rdb初始化错误")
	}
	return nil
}
