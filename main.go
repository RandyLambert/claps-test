package main

import (
	"claps-test/common"
	"claps-test/dao"
	"claps-test/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//读取所有配置文件
	common.InitConfig()

	//连接数据库
	db := dao.InitDB()
	defer db.Close()

	//自动迁移
	/*
	db.Debug().AutoMigrate(&models.Project{})
	db.Debug().AutoMigrate(&models.MemberWallet{})
	db.Debug().AutoMigrate(&models.Repository{})
	db.Debug().AutoMigrate(&models.Transaction{})
	db.Debug().AutoMigrate(&models.Transfer{})
	db.Debug().AutoMigrate(&models.Wallet{})
	 */

	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	//设置session middleware
	store := cookie.NewStore([]byte("claps-test"))
	r.Use(sessions.Sessions("mysession",store))

	r = routers.CollectRoute(r)
	//port := viper.GetString("server.port")
	//if port != "" {
		panic(r.Run(":3001"))
	//}
	//panic(r.Run())

}


