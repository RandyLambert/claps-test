package dao

import (
	"claps-test/model"
	"github.com/jinzhu/gorm"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.MemberWallet{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func UpdateMemberWallet(memberWalletDto *model.MemberWalletDto) (err error) {
	err = db.Debug().Table("member_wallet").Save(memberWalletDto).Error
	return
}

func UpdateMemberWalletBalanceToZeroByUserId(userId int64) (err error) {
	err = db.Debug().Table("member_wallet").Where("user_id = ?", userId).Updates(map[string]interface{}{"balance": "0"}).Error
	return
}

func GetMemberWalletByUserId(userId int64) (memberWalletDtos *[]model.MemberWalletDto, err error) {
	memberWalletDtos = &[]model.MemberWalletDto{}
	err = db.Debug().Table("member_wallet").Where("user_id = ?", userId).Scan(memberWalletDtos).Error
	return
}

func GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(projectId int64, userId int64, botId string, assetId string) (member *model.MemberWalletDto, err error) {
	member = &model.MemberWalletDto{}
	err = db.Debug().Table("member_wallet").Where("project_id=? AND user_id=? AND bot_id=? AND asset_id=?", projectId, userId, botId, assetId).Find(member).Error
	return
}
