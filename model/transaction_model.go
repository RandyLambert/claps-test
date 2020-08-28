package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	Id        string          `json:"id,omitempty" gorm:"type:varchar(50);primary_key;not null;"`
	ProjectId uint32          `json:"project_id,omitempty" gorm:"type:int unsigned;not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);not null"`
	Amount    decimal.Decimal `json:"amount,omitempty" gorm:"type:varchar(128);not null"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	Sender    string `json:"sender,omitempty" gorm:"type:varchar(50);default:null"`
	Receiver  string `json:"receiver,omitempty" gorm:"type:varchar(50);default:null"`
}
