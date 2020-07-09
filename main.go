package main

import (
	"claps-test/common"
	"claps-test/dao"
	"claps-test/models"
	"claps-test/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	common.InitConfig()

	db := dao.InitDB()
	defer db.Close()

	db.Debug().AutoMigrate(&models.Project{})
	db.Debug().AutoMigrate(&models.MemberWallet{})
	db.Debug().AutoMigrate(&models.Repository{})
	db.Debug().AutoMigrate(&models.Transaction{})
	db.Debug().AutoMigrate(&models.Transfer{})
	db.Debug().AutoMigrate(&models.Wallet{})

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


