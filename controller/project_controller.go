package controller

import (
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Projects(ctx *gin.Context) {
	query := &model.PaginationQ{}
	err1 := ctx.ShouldBindQuery(query)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "project query error"), nil)
		return
	}
	projects,number,err := service.ListProjectsByQuery(query)
	if err != nil {
		util.HandleResponse(ctx, err, nil)
		return
	}
	query.Total = number
	resp := &map[string]interface{}{
		"projects":projects,
		"query":query,
	}
	util.HandleResponse(ctx, err, resp)
	//log.Debug(number)
	//util.HandleResponse(ctx,err,projects)
	return
}

/*
通过Id获取项目的详细信息
*/
func ProjectById(ctx *gin.Context) {
	projectId, err1 := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "projectId strconv err"), nil)
		return
	}
	projectInfo, err := service.GetProjectById(projectId)
	util.HandleResponse(ctx, err, projectInfo)
}

func ProjectMembers(ctx *gin.Context) {
	projectId, err1 := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "projectId strconv err"), nil)
		return
	}
	members, err := service.ListMembersByProjectId(projectId)
	util.HandleResponse(ctx, err, members)
}

func ProjectTransactions(ctx *gin.Context) {
	projectId, err1 := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "projectId strconv err"), nil)
		return
	}

	query := &model.PaginationQ{}
	err1 = ctx.ShouldBindQuery(query)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "project query error"), nil)
		return
	}

	transactions, number ,err := service.ListTransactionsByProjectIdAndQuery(projectId,query)
	query.Total = number
	resp := &map[string]interface{}{
		"transactions":transactions,
		"query":query,
	}
	util.HandleResponse(ctx, err, resp)

	return
}

func ProjectSvg(ctx *gin.Context) {
	badge := &model.Badge{}

	if err := ctx.ShouldBindQuery(badge); err != nil {
		err := util.NewErr(nil, util.ErrBadRequest, "没有QUERY值无法请求成功")
		util.HandleResponse(ctx, err, nil)
		return
	}

	err := service.GetProjectBadge(badge)
	util.HandleResponse(ctx, err, nil)
}
