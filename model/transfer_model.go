package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transfer struct {
	//机器人ID
	BotId string `json:"bot_id" gorm:"type:varchar(50);not null"`
	SnapshotId string `json:"snapshot_id" gorm:"type:varchar(50);default null"`
	UserId uint32 `json:"user_id" gorm:"type:int unsigned;not null"`
	TraceId string `json:"trace_id" gorm:"type:varchar(120);primary_key;not null"`
	OpponentId string `json:"opponent_id" gorm:"type:varchar(50);not null"`
	AssetId string `json:"asset_id" gorm:"type:varchar(50);not null"`
	Amount decimal.Decimal `json:"amount" gorm:"type:varchar(128);not null"`
	Memo string `json:"memo" gorm:"type:varchar(120);not null"`
	Status rune `json:"stautus" gorm:"type:enum('0','1');not null"`
	CreatedAt time.Time
}



