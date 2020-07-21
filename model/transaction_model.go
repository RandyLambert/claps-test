package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct{
	Id string `gorm:"type:varchar(50);primary_key;not null;"`
	ProjectId uint32 `gorm:"type:int unsigned;not null"`
	AssetId string `gorm:"type:varchar(50);not null"`
    Amount decimal.Decimal `gorm:"type:varchar(128);not null"`
	CreatedAt time.Time
	Sender string `gorm:"type:varchar(50);default:null"`
	Receiver string `gorm:"type:varchar(50);default:null"`
}

