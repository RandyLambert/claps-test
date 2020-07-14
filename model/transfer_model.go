package model

import "time"

type Transfer struct {
	SnapshotId string `gorm:"type:varchar(50);primary_key;not null"`
	UserId uint32 `gorm:"type:int unsigned;not null"`
	TraceId string `gorm:"type:varchar(50);not null"`
	OpponentId string `gorm:"type:varchar(50);not null"`
	AssetId string `gorm:"type:varchar(50);not null"`
	Amount float64 `gorm:"type:double unsigned;not null"`
	Memo string `gorm:"type:varchar(120);not null"`
	CreatedAt time.Time
}


