package model

type Member struct {
	ProjectId string `gorm:"type:varchar(50);not null;primary_key"`
	UserId    int64  `gorm:"type:bigint;not null;primary_key"`
}
