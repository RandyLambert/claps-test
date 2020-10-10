package service

import (
	"claps-test/model"
	"claps-test/util"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func distributionByMericoAlgorithm(transaction *model.Transaction) {

}

func distributionByCommits(transaction *model.Transaction) {

}

func distributionByChangedLines(transaction *model.Transaction) {

}

//平均分配算法
func distributionByIdenticalAmount(transaction *model.Transaction) {
	//通过项目ID获取所有成员
	members, err := model.USER.ListMembersByProjectId(transaction.ProjectId)
	if err != nil {
		log.Error(err.Error())
		return
	}

	//做除法,如果members等于0上面就返回?
	memberNumbers := decimal.NewFromInt(int64(len(*members)))
	amount := transaction.Amount.Div(memberNumbers)

	if err1 := model.ExecuteTx(func(tx *gorm.DB) error {
		for i := range *members {
			//获得相应的用户钱包
			walletTotal, err := model.MEMBERWALLET.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(transaction.ProjectId, (*members)[i].Id, transaction.Receiver, transaction.AssetId)
			if err != nil {
				log.Error(err.Error())
				return err
			}
			if i == 0 {
				//因为可能会除不尽,所以这里考虑如果出现这种情况,就把除不尽的值转给第一个人
				walletTotal.Total = walletTotal.Total.Add(transaction.Amount.Sub(amount.Mul(memberNumbers.Sub(decimal.NewFromInt(1)))))
				walletTotal.Balance = walletTotal.Balance.Add(transaction.Amount.Sub(amount.Mul(memberNumbers.Sub(decimal.NewFromInt(1)))))
			} else {
				walletTotal.Total = walletTotal.Total.Add(amount)
				walletTotal.Balance = walletTotal.Balance.Add(amount)
			}
			//更新钱包
			err = model.MEMBERWALLET.UpdateMemberWallet(walletTotal)
			if err != nil {
				log.Error(err.Error())
				return err
			}

			err1 := WithdrawNowOrNot(&(*members)[i])
			if err1 != nil {
				log.Error(err1.Error())
				err = errors.New("withdrawalWay等于withdrawByClaps，获得捐赠后直接转账失败")
				return err
			}

		}
		return nil
	}); err1 != nil {
		err := util.NewErr(err1, util.ErrDataBase, "平均分配算法插入提现记录事物出现问题")
		log.Error(err.Error())
		return
	}

}

func WithdrawNowOrNot(member *model.User) (err *util.Err) {
	//判断是否有未完成的提现
	if member.WithdrawalWay == model.WithdrawByClaps {
		if member.MixinId != ""{
			err = IfUnfinishedTransfer(member.MixinId)
			if err != nil {
				return
			}
			//生成transfer记录
			err = DoTransfer(member.Id, member.MixinId)
			if err != nil {
				return
			}
		}
	}
	return
}
