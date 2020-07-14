package main

import (
	"claps-test/common"
	"claps-test/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {


	/*
	初始化配置文件,Mixin,log和DB
	 */
	common.InitConfig()
	common.InitMixin()
	common.InitLog()
	db := common.InitDB()
	defer db.Close()

	common.RegisterType()
	common.Cors()

	log.Debug("debug")
	log.Warningf("Warning")
	log.Error("Error")

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
	serverport := viper.GetString("server.port")
	if serverport != ""{
		panic(r.Run(":"+serverport))
	} else {
		panic(r.Run(":3001"))
	}
}


