package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-gonic/gin"
)

func Projects(ctx *gin.Context) {

	projects, err := service.ListProjectsAll()
	util.HandleResponse(ctx, err, projects)
}

/*
通过Id获取项目的详细信息
 */
func ProjectById(ctx *gin.Context){
	projectInfo, err := service.GetProjectById(ctx.Param("id"))
	util.HandleResponse(ctx, err, projectInfo)
}

func ProjectMembers(ctx *gin.Context) {

	members, err := service.ListMembersByProjectId(ctx.Param("id"))
	util.HandleResponse(ctx, err, members)
}

func ProjectTransactions(ctx *gin.Context) {

	//assetId := ctx.Query("assetId")
	//if assetId == "" {
	//	err := util.NewErr(nil, util.ErrUnauthorized, "没有QUERY值无法请求成功")
	//	util.HandleResponse(ctx, err, nil)
	//	return
	//}

	transactions, err := service.ListTransactionsByProjectId(ctx.Param("id"))
	util.HandleResponse(ctx, err, transactions)
}