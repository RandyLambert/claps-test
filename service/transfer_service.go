package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"github.com/shopspring/decimal"
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
		Status:        '0',
	}

	err1 := dao.InsertTransfer(transfer)

	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"根据分配算法分配之后算出应为没为member分配多少钱,提现记录首次写入数据库失败导致提现失败")
	}
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
