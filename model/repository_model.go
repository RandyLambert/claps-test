package model

import "time"

type Repository struct {
	Id          uint32 `json:"id,omitempty" gorm:"type:int unsigned;primary_key;not null"`
	ProjectId   uint32 `json:"project_id,omitempty" gorm:"type:int unsigned;not null"`
	Type        string `json:"type,omitempty" gorm:"type:enum('GITHUB','BITBUCKET','GITLAB','GIT');not null"`
	Name        string `json:"name,omitempty" gorm:"type:varchar(50);not null"`
	Slug        string `json:"slug,omitempty" gorm:"type:varchar(100);not null"`
	Description string `json:"description,omitempty" gorm:"type:varchar(120);default:null"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
