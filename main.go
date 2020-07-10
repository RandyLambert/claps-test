package main

import (
	"claps-test/common"
	"claps-test/dao"
	"claps-test/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {


	/*
	初始化Mixin,log和DB
	 */
	common.InitMixin()
	common.InitLog()
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
	serverport := viper.GetString("serverport")
	if serverport != ""{
		panic(r.Run(":"+serverport))
	} else {
		panic(r.Run(":3001"))
	}
}


