package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func UpdateMemberWallet(memberWalletDto *model.MemberWalletDto)(err error){
	err = util.DB.Debug().Table("member_wallet").Save(memberWalletDto).Error
	return
}

func InsertMemberWallet(memberWallet *model.MemberWallet)(err error){
	err = util.DB.Create(memberWallet).Error
	return
}

func GetMemeberWalletByUserIdAndAssetId(userId uint32,assetId string)(memberWalletDtos *[]model.MemberWalletDto,err error) {
	memberWalletDtos = &[]model.MemberWalletDto{}
	err = util.DB.Debug().Table("member_wallet").Where("user_id = ? AND asset_id = ?",userId,assetId).Scan(memberWalletDtos).Error
	return
}

func GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(projectId uint32,userId uint32,botId string,assetId string)(member *model.MemberWalletDto,err error){
	member = &model.MemberWalletDto{}
	err = util.DB.Debug().Table("member_wallet").Where("project_id=? AND user_id=? AND bot_id=? AND asset_id=?",projectId,userId,botId,assetId).Find(member).Error
	return
}


