package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func UpdateMemberWallet(memberWallet *model.MemberWallet)(err error){
	err = util.DB.Save(memberWallet).Error
	return
}

func InsertMemberWallet(memberWallet *model.MemberWallet)(err error){
	err = util.DB.Create(memberWallet).Error
	return
}

func GetMemeberWalletByUserIdAndAssetId(userId uint32,assetId string)(memberWallets *[]model.MemberWalletDto,err error) {
	memberWallets = &[]model.MemberWalletDto{}
	err = util.DB.Debug().Table("member_wallet").Where("user_id = ? AND asset_id = ?",userId,assetId).Find(memberWallets).Error
	return
}

func GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(projectId uint32,userId uint32,botId string,assetId string)(member *model.MemberWallet,err error){
	member = &model.MemberWallet{}
	err = util.DB.Where("project_id=? AND user_id=? AND bot_id=? AND asset_id=?",projectId,userId,botId,assetId).Find(member).Error
	return
}


