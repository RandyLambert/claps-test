package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Wallet{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Wallet struct {
	BotId     string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	ProjectId int64           `json:"project_id,omitempty" gorm:"type:bigint;not null"`
	Total     decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
	CreatedAt time.Time       `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`
	UpdatedAt time.Time       `json:"updated_at,omitempty" gorm:"type:timestamp with time zone"`
	SyncedAt  time.Time       `json:"synced_at,omitempty" gorm:"type:timestamp with time zone"`
}

type WalletTotal struct {
	BotId   string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	Total   decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);not null;default:null"`
}

var WALLET *Wallet

func (wallet *Wallet) GetWalletTotalByBotIdAndAssetId(botId string, assetId string) (total *WalletTotal, err error) {
	total = &WalletTotal{}
	err = db.Debug().Table("wallet").Where("bot_id=? AND asset_id=?", botId, assetId).Find(total).Error
	return
}

func (wallet *Wallet) UpdateWalletTotal(walletTotal *WalletTotal) (err error) {
	err = db.Table("wallet").Save(walletTotal).Error
	return
}
