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

//平均分配算法
func distributionByIdenticalAmount(transaction *model.Transaction){
	//通过项目ID获取所有成员
	members,err := dao.ListMembersByProjectId(transaction.ProjectId)
	if err != nil {
		log.Error(err.Error())
		return
	}

	//做除法
	amount := transaction.Amount.Div(decimal.NewFromInt(int64(len(*members))))
	for i := range *members {
		//获得相应的用户钱包
		walletTotal,err := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(transaction.ProjectId,(*members)[i].Id,transaction.Receiver,transaction.AssetId)
		if err != nil {
			log.Error(err.Error())
			return
		}
		walletTotal.Total = walletTotal.Total.Add(amount)
		//更新钱包
		err = dao.UpdateMemberWallet(walletTotal)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

