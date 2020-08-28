package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	BotId     string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	ProjectId uint32          `json:"project_id,omitempty" gorm:"type:int unsigned;not null"`
	Total     decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
	CreatedAt time.Time		  `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	SyncedAt  time.Time       `json:"synced_at,omitempty"`
}

type WalletTotal struct {
	BotId   string          `json:"bot_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	AssetId string          `json:"asset_id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	Total   decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);not null;default:null"`
}
