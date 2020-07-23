package model

import (
	"github.com/shopspring/decimal"
	"time"
)

const  (
	INITIAL = "0"
	UNFINISHED = "1"
	FINISHED = "2"
)

type Transfer struct {
	//机器人ID
	TransactionId string `json:"transaction_id" gorm:"type:varchar(50);not null"`
	BotId string `json:"bot_id" gorm:"type:varchar(50);not null"`
	SnapshotId string `json:"snapshot_id" gorm:"type:varchar(50);default null"`
	UserId uint32 `json:"user_id" gorm:"type:int unsigned;not null"`
	TraceId string `json:"trace_id" gorm:"type:varchar(120);not null;primary_key"`
	AssetId string `json:"asset_id" gorm:"type:varchar(50);not null"`
	Amount decimal.Decimal `json:"amount" gorm:"type:varchar(128);not null"`
	Memo string `json:"memo" gorm:"type:varchar(120);not null"`
	Status string `json:"stautus" gorm:"type:enum('0','1','2');not null"`
	//0按照分配算法生成之后的状态 1用户点击提现后(提现操作未完成) 2机器人完成提现
	CreatedAt time.Time
}



