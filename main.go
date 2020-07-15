package main

import (
	"claps-test/model"
	"claps-test/util"
	"claps-test/router"
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
	util.InitConfig()
	util.InitMixin()
	util.InitLog()
	db := util.InitDB()
	defer db.Close()

	util.RegisterType()
	util.Cors()

	log.Debug("debug")
	log.Warning("Warning")
	log.Error("Error")

	//自动迁移
	db.Debug().AutoMigrate(&model.Project{})
	db.Debug().AutoMigrate(&model.MemberWallet{})
	db.Debug().AutoMigrate(&model.Repository{})
	db.Debug().AutoMigrate(&model.Transaction{})
	db.Debug().AutoMigrate(&model.Transfer{})
	db.Debug().AutoMigrate(&model.Wallet{})
	db.Debug().AutoMigrate(&model.User{})

	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	//设置session middleware
	store := cookie.NewStore([]byte("claps-test"))
	r.Use(sessions.Sessions("mysession",store))

	r = router.CollectRoute(r)
	serverport := viper.GetString("server.port")
	if serverport != ""{
		panic(r.Run(":"+serverport))
	} else {
		panic(r.Run(":3001"))
	}
}


