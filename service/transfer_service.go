package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func InsertTransfer(botId, assetID , memo string, amount decimal.Decimal,userId uint32)(err error){
	transfer := &model.Transfer{
		BotId:         botId,
		UserId:        userId,
		TraceId:       string(userId)+assetID,
		//TransactionId: TransactionID,
		AssetId:       assetID,
		Amount:        amount,
		Memo:          memo,
		Status:        model.UNFINISHED,
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

func DoTransfer(userId uint32,assetId string) (err *util.Err) {

	memberWallets,err1 := dao.GetMemeberWalletByUserIdAndAssetId(userId,assetId)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取用户钱包失败导致提现失败")
		return
	}

	for i := range *memberWallets {
		mixinId,err1 := dao.GetMixinIdByUserId(userId)
		if err1 != nil {
			err = util.NewErr(err1,util.ErrDataBase,"获取用户MixinId失败导致提现失败")
			return
		}
		err1 = InsertTransfer(mixinId.MixinId,assetId,"test",(*memberWallets)[i].Balance,userId)
		if err1 != nil {
			err = util.NewErr(err1,util.ErrDataBase,"插入捐赠记录失败")
			return
		}
		(*memberWallets)[i].Balance = decimal.Zero
		err1 = dao.UpdateMemberWallet(&(*memberWallets)[i])
		if err1 != nil {
			err = util.NewErr(err1,util.ErrDataBase,"更新用户钱包可提现值导致提现失败")
			return
		}
	}

	return
}
