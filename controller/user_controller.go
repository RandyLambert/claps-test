package controller

import (
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	"github.com/shopspring/decimal"
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
	projects,err := service.ListProjectsByUserId(uint32(*user.(github.User).ID))
	if err != nil {
		util.HandleResponse(ctx,err,resp)
	}

	log.Debug("获得的projects是",projects)
	resp["emails"] = emails
	resp["projects"] = projects
	util.HandleResponse(ctx,err,resp)

}

//获取捐赠记录,就是个人的提现记录
func Transfer(ctx *gin.Context) {
	resp := make(map[string]interface{})

	assetId := ctx.Query("assetId")
	if assetId == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"没有币种参数")
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug("assetId = ",assetId)

	//从transaction表中获取该用户的所有捐赠记录
}

//获取用户钱包中所有币种的余额
func UserAssets(ctx *gin.Context){

	var err *util.Err
	resp := make(map[string]interface{})
	session := sessions.Default(ctx)
	//user不为空
	user := session.Get("user")
	userId := user.(github.User).ID

	//获得所有币的信息
	assets,err := service.ListAssetsAllByDB()
	if err != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug(*assets)

	//查询用户钱包,获得相应的余额,添加到币信息的后面
	err2,dto := service.GetUserBalanceByAllAssets(uint32(*userId),assets)
	if err2 != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}

	log.Debug(*dto)
	resp["assets"] = dto
	util.HandleResponse(ctx,err,resp)
}


//从transaction中读取关于自己项目的所有捐赠
func UserTransactions(ctx *gin.Context){
	resp := make(map[string]interface{})

	assetId := ctx.Query("assetId")
	if assetId == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"没有币种参数")
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug("assetId = ",assetId)

	//从transfer表中获取该用户的所有捐赠记录
}

//读取某种货币的交易记录,读取transfer里面的记录/
func UserTransfer(ctx *gin.Context) {
	resp := make(map[string]interface{})
	session := sessions.Default(ctx)
	userId := uint32(*(session.Get("user").(github.User).ID))

	assetId := ctx.Query("assetId")
	if assetId == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"没有币种参数")
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug("assetId = ",assetId)

	//从transfer表中获取该用户的所有捐赠记录
	transfers,err := service.GetTransferByUserIdAndAssetId(userId,assetId)
	resp["transfers"] = transfers
	util.HandleResponse(ctx,err,resp)
}

//获取某用户的所有的受捐赠记录的汇总
func UserDonation(ctx *gin.Context) {
	resp := make(map[string]interface{})
	session := sessions.Default(ctx)
	userId := uint32(*(session.Get("user").(github.User)).ID)

	//读取所有的member_wallet表然后汇总
	//获得所有币的信息
	assets,err := service.ListAssetsAllByDB()
	if err != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}
	log.Debug(*assets)

	//查询用户钱包,获得相应的余额,添加到币信息的后面
	err2,dto := service.GetUserBalanceByAllAssets(userId,assets)
	if err2 != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}

	//便利dto然后求和
	var sum decimal.Decimal
	var patrons uint32

	for i := range *dto{
		sum = sum.Add((*dto)[i].Balance)
	}

	//从project里面寻找patrons然后求和
	patrons,err3 := service.GetPatronsByUserId(userId)
	if err3 != nil{
		util.HandleResponse(ctx,err3,resp)
		return
	}

	resp["total"] = sum
	resp["patrons"] = patrons

	util.HandleResponse(ctx,nil,resp)
}

//用户提现某种货币,把表中的status由0变为1
func UserWithdraw(ctx *gin.Context) {

	//获取币种
	assetId := ctx.Query("assetId")
	if assetId == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"请求路由无assetId参数")
		util.HandleResponse(ctx,err,nil)
		return
	}

	//获取userId
	session := sessions.Default(ctx)
	userId := uint32(*session.Get("user").(github.User).ID)

	//判断是否有未完成的提现
	 err3 := service.IfUnfinishedTransfer(userId,assetId)
	 if err3 != nil{
	 	util.HandleResponse(ctx,err3,nil)
		 return
	 }

	//找到相应的币种的所有transfer然后改变status
	err2 := service.UpdateTransferStatusByAssetIdAndUserId(userId,assetId,model.UNFINISHED)
	if err2 != nil{
		util.HandleResponse(ctx,err2,nil)
		return
	}
	util.HandleResponse(ctx,nil,nil)

	//等协程完成转账
}


