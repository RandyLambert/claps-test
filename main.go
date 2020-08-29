package main

import (
	"claps-test/dao"
	"claps-test/router"
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	if rediserr := util.InitClient();rediserr != nil{
		log.Error(rediserr)
	}

	db,_ := dao.InitDB()
	if db != nil {
		defer db.Close()
	}

	//自动迁移
	if multierror := dao.Migrate(); multierror != nil{
		log.Error(multierror)
	}

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
