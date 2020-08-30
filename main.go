package main

import (
	"claps-test/dao"
	"claps-test/router"
	"claps-test/service"
	"claps-test/util"
	"flag"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
func init() {
	/*
		初始化配置文件,Mixin,log,DB和cache
	*/
	util.InitConfig()
	util.InitMixin()
	if err := util.InitClient();err != nil{
		log.Error(err)
	}
}

func main() {

	cmd := flag.String("cmd", "", "process identity")
	flag.Parse()

	db, _ := dao.InitDB()
	if db != nil {
		defer db.Close()
	}

	switch *cmd {
	case "migrate", "setdb":
		if multierror := dao.Migrate(); multierror != nil {
			log.Error(multierror)
		}
	default:
		//定期更新数据库snapshot信息
		go service.SyncSnapshots()
		//定期更新数据库asset信息
		go service.SyncAssets()
		//定期进行提现操作,并更改数据库
		go service.SyncTransfer()

		util.RegisterType()
		util.Cors()

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
}
