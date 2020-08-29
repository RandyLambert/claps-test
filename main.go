package main

import (
	"claps-test/model"
	"claps-test/router"
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	if db != nil {
		defer db.Close()
	}

	db.Save(&model.Property{
		Key:   "last_snapshot_id",
		Value: "1",
	})
	//自动迁移
	db.Debug().AutoMigrate(&model.Project{})
	db.Debug().AutoMigrate(&model.MemberWallet{})
	db.Debug().AutoMigrate(&model.Repository{})
	db.Debug().AutoMigrate(&model.Transaction{})
	db.Debug().AutoMigrate(&model.Transfer{})
	db.Debug().AutoMigrate(&model.Wallet{})
	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Member{})
	db.Debug().AutoMigrate(&model.Property{})
	db.Debug().AutoMigrate(&model.Bot{})
	db.Debug().AutoMigrate(&model.Asset{})

	util.RegisterType()
	util.Cors()
	//定期更新数据库snapshot信息
	go service.SyncSnapshots()
	//定期更新数据库asset信息
	go service.SyncAssets()
	//定期进行提现操作,并更改数据库
	go service.SyncTransfer()

	r := gin.Default()

	//设置session middleware
	store := cookie.NewStore([]byte("claps-test"))
	r.Use(sessions.Sessions("mysession", store))

	r = router.CollectRoute(r)
	serverport := viper.GetString("server.port")
	if serverport != "" {
		panic(r.Run(":" + serverport))
	} else {
		panic(r.Run(":3001"))
	}
}
