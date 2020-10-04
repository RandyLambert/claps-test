package model

import "github.com/shopspring/decimal"

type Fiat struct {
	Code string          `json:"code,omitempty" gorm:"type:varchar(25);primary_key;not null"`
	Rate decimal.Decimal `json:"rate,omitempty" gorm:"type:varchar(128);not null"`
}
