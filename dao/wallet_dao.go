package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func GetWalletByBotIdAndAssetId(botId string,assetId string)(wallet *model.Wallet,err error){
	wallet = &model.Wallet{}
	err = util.DB.Where("bot_id=? AND asset_id=?",botId,assetId).Find(wallet).Error
	return
}

func UpdateWallet(wallet *model.Wallet)(err error){
	err = util.DB.Save(wallet).Error
	return
}

//通过项目Id获取项目的Total?
func GetWalletTotalByProjectId(projectId uint32)(total *[]model.WalletTotal,err error){
	total = &[]model.WalletTotal{}
	err = util.DB.Debug().Table("wallet").Select("project_id,asset_id,total").Where("project_id=?",projectId).Scan(total).Error
	return
}

