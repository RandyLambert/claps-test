package models

import "time"

type Transaction struct{
	Id string `gorm:"type:varchar(50);primary_key;not null;"`
	ProjectId uint32 `gorm:"type:int unsigned;not null"`
	BotId string `gorm:"type:varchar(50);not null"`
	AssetId string `gorm:"type:varchar(50);not null"`
    Amount float64 `gorm:"type:double unsigned;not null"`
	CreatedAt time.Time
	Sender string `gorm:"type:varchar(50);default:null"`
	Receiver string `gorm:"type:varchar(50);default:null"`
}

