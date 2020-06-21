package models

import "time"

type Repository struct {
	Id uint32 `gorm:"type:int unsigned;primary_key;not null;unique_index:id_UNIQUE"`
	ProjectId uint32 `gorm:"type:int unsigned;not null"`
	Type string `gorm:"type:enum('GITHUB','BITBUCKET','GITLAB','GIT');not null"`
	Name string `gorm:"type:varchar(50);not null"`
	Slug string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(120);default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


