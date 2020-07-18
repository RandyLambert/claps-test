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

