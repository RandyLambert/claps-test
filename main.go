package main

import (
	"claps-test/common"
	"claps-test/dao"
	"claps-test/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	_"github.com/go-sql-driver/mysql"
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

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}


