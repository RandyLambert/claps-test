package controllers

import (
	"claps-test/dao"
	"claps-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Progects(ctx *gin.Context){
	db := dao.GetDB()
	var projects []models.Project
	db.Debug().Find(&projects)
	ctx.JSON(http.StatusOK,projects)
	//var bots []models.Bot
	//db.Debug().Find(&bots)
	//ctx.JSON(http.StatusOK,bots)

}


func Project(ctx *gin.Context){

	db := dao.GetDB()
	name := ctx.Param("name")
	var project models.Project
	db.Debug().Where("name=?",name).Find(&project)
	ctx.JSON(http.StatusOK,project)

	//var bots []models.Bot
	//db.Debug().Find(&bots)
	//ctx.JSON(http.StatusOK,bots)

}

func ProgectMembers(ctx *gin.Context){

	db := dao.GetDB()
	name := ctx.Param("name")
	var users []models.User
	db.Debug().Joins("INNER JOIN member ON member.user_id = user.id").Joins("INNER JOIN project ON project.name=?",name).Where("project.id=?","member.project_id").Find(&users)
	ctx.JSON(http.StatusOK,users)
}

func ProgectTransactions(ctx *gin.Context){
	db := dao.GetDB()
	name := ctx.Param("name")
	assetId := ctx.Query("assetId")
	var transactions []models.Transaction
	db.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(&transactions)
	ctx.JSON(http.StatusOK,transactions)
}
