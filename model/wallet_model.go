package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	BotId     string          `gorm:"type:varchar(50);primary_key;not null"`
	AssetId   string          `gorm:"type:varchar(50);primary_key;not null"`
	ProjectId uint32          `gorm:"type:int unsigned;not null"`
	Total     decimal.Decimal `gorm:"type:varchar(128);default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	SyncedAt  time.Time
}

type WalletTotal struct {
	BotId   string          `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string          `gorm:"type:varchar(50);primary_key;not null"`
	Total   decimal.Decimal `gorm:"type:varchar(128);not null;default:null"`
}
