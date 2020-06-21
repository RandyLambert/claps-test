package models

import "time"

type Project struct {
	Id uint32 `gorm:"type:int unsigned;primary_key;not null;unique_index:id_UNIQUE"`
	Name string `gorm:"type:varchar(50);not null;unique_index:name_UNIQUE"`
    DisplayName string `gorm:"type:varchar(50);default:null"`
	Description string `gorm:"type:varchar(120);default:null"`
	AvatarUrl string `gorm:"type:varchar(100);default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


