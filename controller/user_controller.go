package controller

import (
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
功能:返回用户的邮箱数组和项目数组
说明:执行该函数,一定登录了github
*/
func UserProfile(ctx *gin.Context) {
	var err *util.Err
	resp := make(map[string]interface{})

	//获取github信息
	mcache := util.MCache{}
	uid := ctx.GetString(util.UID)
	err1 := util.Rdb.Get(uid, &mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error."), nil)
		return
	}
	log.Debug("mcache = ", mcache)

	//根据userId获取所有project信息,Total和Patrons字段添加
	projects, err := service.ListProjectsByUserId(*mcache.Github.ID)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	resp["emails"] = mcache.GithubEmails
	resp["projects"] = projects
	util.HandleResponse(ctx, err, resp)
}

/*
功能:获取用户钱包所有币种的余额
说明:此时已经登录github,不需要绑定mixin
*/
func UserAssets(ctx *gin.Context) {

	var err *util.Err
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "cache get uid error"), resp)
		return
	}
	uid := val.(string)

	//从redis获取cache
	mcache := &util.MCache{}
	err1 := util.Rdb.Get(uid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), resp)
		return
	}

	//获得所有币的信息
	assets, err := service.ListAssetsAllByDB()
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	//查询用户钱包,获得相应的余额,添加到币信息的后面
	err2, dto := service.GetBalanceAndTotalByUserIdAndAssets(*mcache.Github.ID, assets)
	if err2 != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	resp["assets"] = dto
	util.HandleResponse(ctx, err, resp)
}

/*
功能:获取某币种的交易记录,从transfer表里面读取数据
说明:用户一定登录了github和mixin,中间件保证
*/
func UserTransfer(ctx *gin.Context) {
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), resp)
		return
	}
	uid := val.(string)

	mcache := &util.MCache{}
	err1 := util.Rdb.Get(uid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), resp)
		return
	}

	query := &model.PaginationQ{}
	err1 = ctx.ShouldBindQuery(query)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrBadRequest, "transfer query error"), nil)
		return
	}

	//从transfer表中获取该用户的所有捐赠记录
	transfers, number, err := service.ListTransfersByProjectIdAndQuery(mcache.MixinId, query)
	if err != nil {
		util.HandleResponse(ctx, err, nil)
		return
	}
	query.Total = number
	resp["transfers"] = transfers
	resp["query"] = query

	util.HandleResponse(ctx, err, resp)
}

//获取某用户的所有的受捐赠记录的汇总
/*
功能:请求获得某个用户的捐赠信息的汇总,包括总金额和捐赠人数
说明:不需要绑定mixin
*/
func UserDonation(ctx *gin.Context) {
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), resp)
		return
	}
	uid := val.(string)

	mcache := &util.MCache{}
	err1 := util.Rdb.Get(uid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), resp)
		return
	}

	//读取所有的member_wallet表然后汇总
	//获得所有币的信息
	assets, err := service.ListAssetsAllByDB()
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}
	log.Debug(*assets)

	//查询用户钱包,获得相应的余额,添加到币信息的后面
	err2, total, balance := service.GetBalanceAndTotalToUSDByUserId(*mcache.Github.ID, assets)
	if err2 != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	//从project里面寻找Donations
	donations, err3 := service.SumProjectDonationsByUserId(*mcache.Github.ID)
	if err3 != nil {
		util.HandleResponse(ctx, err3, resp)
		return
	}

	resp["total"] = total
	resp["balance"] = balance
	resp["donations"] = donations

	util.HandleResponse(ctx, nil, resp)
}

/*
功能:用户提现某种货币,把表中的status由0变为1
说明:已经中间件验证绑定了mixin
*/
func UserWithdraw(ctx *gin.Context) {
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), resp)
		return
	}
	uid := val.(string)

	mcache := &util.MCache{}
	err1 := util.Rdb.Get(uid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), resp)
		return
	}

	//已经绑定mixin,直接从缓存中取
	mixinId := mcache.MixinId

	//判断是否有未完成的提现
	err3 := service.IfUnfinishedTransfer(mixinId)
	if err3 != nil {
		util.HandleResponse(ctx, err3, nil)
		return
	}

	//生成transfer记录
	err2 := service.DoTransfer(*mcache.Github.ID, mixinId)
	if err2 != nil {
		util.HandleResponse(ctx, err2, nil)
		return
	}
	util.HandleResponse(ctx, nil, nil)

	//等协程完成转账
}

func UserWithdrawalWay(ctx *gin.Context) {
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), resp)
		return
	}
	uid := val.(string)

	mcache := &util.MCache{}
	err1 := util.Rdb.Get(uid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), resp)
		return
	}

	withdrawalWay := ctx.DefaultPostForm("withdrawal_way", model.WithdrawByClaps)
	if withdrawalWay != model.WithdrawByClaps && withdrawalWay != model.WithdrawByUser {
		util.HandleResponse(ctx, util.NewErr(errors.New("该请求post值不属于withdrawByClaps或者withdrawByUser"),
			util.ErrBadRequest, "该请求post值不属于withdrawByClaps或者withdrawByUser"), resp)
		return
	}

	//更新withdrawalWay
	err2 := service.UpdateUserWithdrawalWay(*mcache.Github.ID, withdrawalWay)
	if err2 != nil {
		util.HandleResponse(ctx, err2, nil)
		return
	}
	util.HandleResponse(ctx, nil, nil)
}
