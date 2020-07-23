package service

import (
	"claps-test/dao"
	"claps-test/model"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func distributionByPersperAlgorithm(transaction *model.Transaction){

}

func distributionByCommits(transaction *model.Transaction){

}

func distributionByChangedLines(transaction *model.Transaction){

}

func distributionByIdenticalAmount(transaction *model.Transaction){
	members,err := dao.ListMembersByProjectId(transaction.ProjectId)
	if err != nil {
		log.Error(err.Error())
		return
	}
	amount := transaction.Amount.Div(decimal.NewFromInt(int64(len(*members))))
	for i := range *members {
		walletTotal,err := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(transaction.ProjectId,(*members)[i].Id,transaction.Receiver,transaction.AssetId)
		if err != nil {
			log.Error(err.Error())
			return
		}
		walletTotal.Total = walletTotal.Total.Add(amount)
		err = dao.UpdateMemberWallet(walletTotal)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

