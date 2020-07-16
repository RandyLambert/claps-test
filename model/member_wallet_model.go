package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type MemberWallet struct {
	ProjectId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	UserId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	BotId string `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string `gorm:"type:varchar(50);primary_key;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Total decimal.Decimal `gorm:"type:varchar(50);default:null"`
	Balance decimal.Decimal `gorm:"type:varchar(50);default:null"`
}


