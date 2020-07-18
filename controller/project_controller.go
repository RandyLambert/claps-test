package controller

import (
	"claps-test/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Projects(ctx *gin.Context){
	//
	projects,err := service.ListProjectsAll()
	if err !=nil {
		log.Error("Projects",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}
	ctx.JSON(http.StatusOK,projects)

}


func Project(ctx *gin.Context){
	projectInfo,err := service.GetProjectByName(ctx.Param("name"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}else{
		ctx.JSON(http.StatusOK,projectInfo)
	}

	//var bots []model.Bot
	//db.Debug().Find(&bots)
	//ctx.JSON(http.StatusOK,bots)

}

func ProjectMembers(ctx *gin.Context){

	members,err := service.ListMembersByProjectName(ctx.Param("name"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}
	ctx.JSON(http.StatusOK,members)
}

func ProjectTransactions(ctx *gin.Context){
	transactions,err := service.ListTransactionsByProjectNameAndAssetId(ctx.Param("name"),ctx.Query("assetId"))
	if err != nil {
		log.Error("Project",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err,
		})
	}else {
		ctx.JSON(http.StatusOK,transactions)
	}
}
