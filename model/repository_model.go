package model

import "time"

type Repository struct {
	Id          uint32 `json:"id" gorm:"type:int unsigned;primary_key;not null"`
	ProjectId   uint32 `json:"project_id" gorm:"type:int unsigned;not null"`
	Type        string `json:"type" gorm:"type:enum('GITHUB','BITBUCKET','GITLAB','GIT');not null"`
	Name        string `json:"name" gorm:"type:varchar(50);not null"`
	Slug        string `json:"slug" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:varchar(120);default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
