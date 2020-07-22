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
		err = InsertTransfer(transaction.Receiver,transaction.AssetId,transaction.Id,"test",amount,(*members)[i].Id)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

