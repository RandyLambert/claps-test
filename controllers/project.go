package controllers

import (
	"claps-test/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Projects(ctx *gin.Context){
	//
	projects,err := service.GetProjects()
	if err !=nil {
		log.Error("Projects",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}
	ctx.JSON(http.StatusOK,projects)

}


func Project(ctx *gin.Context){
	projectInfo,err := service.GetProject(ctx.Param("name"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}else{
		ctx.JSON(http.StatusOK,projectInfo)
	}

	//var bots []models.Bot
	//db.Debug().Find(&bots)
	//ctx.JSON(http.StatusOK,bots)

}

func ProjectMembers(ctx *gin.Context){

	members,err := service.GetProjectMembers(ctx.Param("name"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}
	ctx.JSON(http.StatusOK,members)
}

func ProjectTransactions(ctx *gin.Context){
	transactions,err := service.GetProjectTransactions(ctx.Param("name"),ctx.Query("assetId"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}
	ctx.JSON(http.StatusOK,transactions)
}
