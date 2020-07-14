package model

import "time"

type Wallet struct {
	BotId string `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string `gorm:"type:varchar(50);primary_key;not null"`
	ProjectId uint32 `gorm:"type:int unsigned;not null"`
	Total float64 `gorm:"type:double unsigned;not null;default:0"`
	Balance float64 `gorm:"type:double unsigned;not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	SyncedAt time.Time

}

type ProjectTotal struct {
	Total float64 `gorm:"type:double unsigned;not null;default:0"`
	AssetId string `gorm:"type:double unsigned;not null;default:0"`
}