package models

import "time"

type MemberWallet struct {
	ProjectId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	UserId uint32 `gorm:"type:int unsigned;primary_key;not null"`
	BotId string `gorm:"type:varchar(50);primary_key;not null"`
	AssetId string `gorm:"type:varchar(50);primary_key;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Total float64 `gorm:"type:double unsigned;not null;default:0"`
	Balance float64 `gorm:"type:double unsigned;not null;default:0"`
}


