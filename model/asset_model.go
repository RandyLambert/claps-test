package model

import "github.com/shopspring/decimal"

type Asset struct {
	AssetId  string          `json:"asset_id" gorm:"type:varchar(36);primary_key;not null;"`
	Symbol   string          `json:"symbol" gorm:"type:varchar(512);not null"`
	Name     string          `json:"name" gorm:"type:varchar(512);not null"`
	IconUrl  string          `json:"icon_url" gorm:"type:varchar(1024);not null"`
	PriceBtc decimal.Decimal `json:"price_btc" gorm:"type:varchar(128);not null"`
	PriceUsd decimal.Decimal `json:"price_usd" gorm:"type:varchar(128);not null"`
}

type AssetIdToPriceUsd struct {
	AssetId  string          `json:"asset_id" gorm:"type:varchar(36);primary_key;not null;"`
	PriceUsd decimal.Decimal `json:"price_usd" gorm:"type:varchar(128);not null"`
}
