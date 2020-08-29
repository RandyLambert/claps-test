package util

import (
	"github.com/gin-contrib/cache/persistence"
	"time"
)

// 声明一个全局的rdb变量
var Rdb *persistence.InMemoryStore

//jwt的过期时间
const TokenExpireDuration = time.Minute* 2

// 初始化连接
func InitClient() (err error) {
	Rdb = persistence.NewInMemoryStore(TokenExpireDuration)

	/*
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("REDIS_ADDR"),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       config.GetInt("REDIS_DB"),  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	 */
	return nil
}