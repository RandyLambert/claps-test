package model

type ProjectToMericoGroup struct {
	ProjectId     int64  `gorm:"type:bigint;primary_key;not null"`
	MericoGroupId string `gorm:"type:varchar(50);primary_key;not null"`
}
