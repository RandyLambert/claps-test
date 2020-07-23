package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func InsertTransfer(botId, assetID ,TransactionID, memo string, amount decimal.Decimal,userId uint32)(err error){
	transfer := &model.Transfer{
		BotId:         botId,
		UserId:        userId,
		TraceId:       string(userId)+TransactionID+assetID,
		TransactionId: TransactionID,
		AssetId:       assetID,
		Amount:        amount,
		Memo:          memo,
		Status:        model.INITIAL,
	}

	err1 := dao.InsertTransfer(transfer)

	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"根据分配算法分配之后算出应为没为member分配多少钱,提现记录首次写入数据库失败导致提现失败")
	}
	return
}

//跟新数据库中的transger表中的status为1
func UpdateTransferStatusByAssetIdAndUserId(userId uint32,assetId string,status string)(err *util.Err){
	err1 := dao.UpdateTransferSatusByUserIdAndAssetId(userId,assetId,status)
	if err1 != nil{
		err = util.NewErr(err1,util.ErrDataBase,"更新数据库transfer状态出错")
	}
	return
}

//判断某种币是否有未完成的提现操作,err非nil标有有未完成,err=nil表示没有未完成
func IfUnfinishedTransfer(userId uint32,assetId string) (err *util.Err) {
	count,err1 := dao.CountUnfinishedTransfer(userId,assetId)
	if err1 != nil{
		err = util.NewErr(err1,util.ErrDataBase,"数据库查询Unfinished出错")
		return
	}
	if count != 0{
		err = util.NewErr(errors.New("有未完成的提现操作"),util.ErrForbidden,"有未完成的提现操作")
		return
	}

	log.Debug(count)
	return
}

func DoTransfer(projectId,userId uint32,botId,assetId string,amount decimal.Decimal) (err *util.Err) {
	//这里的memberwallet,是通过外部获取的,业务逻辑不是这样,暂时这么写
	memberWallet,err1 := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(projectId,userId,botId,assetId)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取用户钱包失败导致提现失败")
		return
	}

	memberWallet.Balance = memberWallet.Balance.Sub(amount)

	err1 = dao.UpdateMemberWallet(memberWallet)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"更新用户钱包可提现值导致提现失败")
	}
	return
}
