package main

import (
	"claps-test/common"
	"claps-test/dao"
	"claps-test/models"
	"claps-test/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
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
	r.LoadHTMLFiles("views/hello.html")
	r = routers.CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}


