package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type MemberWallet struct {
	ProjectId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	//user表的Id
	UserId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	BotId string `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string `gorm:"type:varchar(50);primary_key;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Total decimal.Decimal `gorm:"type:varchar(128);default:null"`
	Balance decimal.Decimal `gorm:"type:varchar(128);default:null"`
}

type MemberWalletDto struct {
	AssetId string `json:"asset_id" gorm:"type:varchar(50);primary_key;not null"`
	Balance decimal.Decimal `json:"balance"  gorm:"type:varchar(128);default:null"`
}
