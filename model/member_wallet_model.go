package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&MemberWallet{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type MemberWallet struct {
	ProjectId int64 		  `json:"project_id,omitempty" gorm:"type:bigint;primary_key;not null"`
	//user表的Id
	UserId    int64           `json:"user_id,omitempty" gorm:"type:bigint;primary_key;not null"`
	BotId     string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	CreatedAt time.Time       `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`
	UpdatedAt time.Time       `json:"updated_at,omitempty" gorm:"type:timestamp with time zone"`
	Total     decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
	Balance   decimal.Decimal `json:"balance,omitempty" gorm:"type:varchar(128);default:null"`
}

type MemberWalletDto struct {
	ProjectId int64           `json:"project_id,omitempty" gorm:"type:bigint;primary_key;not null"`
	UserId    int64           `json:"user_id,omitempty" gorm:"type:bigint;primary_key;not null"`
	BotId     string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	Total     decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
	Balance   decimal.Decimal `json:"balance,omitempty" gorm:"type:varchar(128);default:null"`
}

var MEMBERWALLET *MemberWallet

func (MemberWallet *MemberWallet) UpdateMemberWallet(memberWalletDto *MemberWalletDto) (err error) {
	err = db.Debug().Table("member_wallet").Save(memberWalletDto).Error
	return
}

func (MemberWallet *MemberWallet) UpdateMemberWalletBalanceToZeroByUserId(userId int64) (err error) {
	err = db.Debug().Table("member_wallet").Where("user_id = ?", userId).Updates(map[string]interface{}{"balance": "0"}).Error
	return
}

func (MemberWallet *MemberWallet) GetMemberWalletByUserId(userId int64) (memberWalletDtos *[]MemberWalletDto, err error) {
	memberWalletDtos = &[]MemberWalletDto{}
	err = db.Debug().Table("member_wallet").Where("user_id = ?", userId).Scan(memberWalletDtos).Error
	return
}

func (MemberWallet *MemberWallet) GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(projectId int64, userId int64, botId string, assetId string) (member *MemberWalletDto, err error) {
	member = &MemberWalletDto{}
	err = db.Debug().Table("member_wallet").Where("project_id=? AND user_id=? AND bot_id=? AND asset_id=?", projectId, userId, botId, assetId).Find(member).Error
	return
}