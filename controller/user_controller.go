package controller

import (
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
)


//获取用户的邮箱和项目
func UserProfile(ctx *gin.Context){
	var err *util.Err
	resp := make(map[string]interface{})

	session := sessions.Default(ctx)
	//从session中获取user和githubToken信息
	user := session.Get("user")
	githubToken := session.Get("githubToken")

	if githubToken == nil{
		//未授权
		err = util.NewErr(errors.New("未授权"),util.ErrUnauthorized,"未授权")
		util.HandleResponse(ctx,err,nil)
	}

	token := githubToken.(string)
	log.Debug("profile中的user:",user)
	log.Debug("profile中的githubToken:",token)

	//获取email信息
	emails,err := service.ListEmailsByToken(token)
	//如果因为超时出错,重新请求

	if err != nil {
		util.HandleResponse(ctx,err,resp)
		return
	}

	//根据userId获取所有project信息,Total和Patrons字段添加
	projects,err := service.ListProjectsByUserId(*user.(github.User).ID)
	if err != nil {
		util.HandleResponse(ctx,err,resp)
	}

	log.Debug("获得的projects是",projects)
	resp["emails"] = emails
	resp["projects"] = projects
	util.HandleResponse(ctx,err,resp)

}

//获取用户钱包中所有币种的余额
func UserAssets(ctx *gin.Context){

	var err *util.Err
	resp := make(map[string]interface{})

	//请求
	assets,err := service.ListAssets()

	log.Debug(assets,err)

	//
	util.HandleResponse(ctx,err,resp)
}


//读取交易记录
func UserTransactions(ctx *gin.Context){
	resp := make(map[string]interface{})

	assetId := ctx.Query("assetId")
	if assetId == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"没有币种参数")
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug("assetId = ",assetId)


}





