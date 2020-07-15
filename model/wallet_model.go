package model

import "time"

type Wallet struct {
	BotId string `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string `gorm:"type:varchar(50);primary_key;not null"`
	ProjectId uint32 `gorm:"type:int unsigned;not null"`
	Total string `gorm:"type:varchar(50);not null;default:null"`
	Balance string `gorm:"type:varchar(50);not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	SyncedAt time.Time

}

type ProjectTotal struct {
	Total string `gorm:"type:varchar(50);not null;default:null"`
	AssetId string `gorm:"type:double unsigned;not null;default:0"`
}