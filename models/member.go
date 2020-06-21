package models

type Member struct {
	ProjectId uint32 `gorm:"type:int unsigned;not null;primary_key"`
	UserId uint32 `gorm:"type:int unsigned;not null;primary_key"`
}

