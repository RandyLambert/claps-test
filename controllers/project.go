package controllers

import (
	"claps-test/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Projects(ctx *gin.Context){
	//todo 数组前缀items
	projects := dao.GetProjects()
	ctx.JSON(http.StatusOK,projects)

}


func Project(ctx *gin.Context){
    //todo github star
	project := dao.GetProject(ctx.Param("name"))
	//projectinfo :=
	ctx.JSON(http.StatusOK,project)

	//var bots []models.Bot
	//db.Debug().Find(&bots)
	//ctx.JSON(http.StatusOK,bots)

}

func ProjectMembers(ctx *gin.Context){
	//todo mambers格式
	members := dao.GetProjectMembers(ctx.Param("name"))
	users := map[string]interface{}{
		"members":members,
	}

	ctx.JSON(http.StatusOK,users)
}

func ProjectTransactions(ctx *gin.Context){

	transactions := dao.GetProjectTransactions(ctx.Param("name"),ctx.Query("assetId"))
	ctx.JSON(http.StatusOK,transactions)
}
