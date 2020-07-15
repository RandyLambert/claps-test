package model

import "time"

type Project struct {
	Id uint32 `json:"id" gorm:"type:int unsigned;primary_key;not null;unique_index:id_UNIQUE"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique_index:name_UNIQUE"`
    DisplayName string `json:"display_name" gorm:"type:varchar(50);default:null"`
	Description string `json:"description" gorm:"type:varchar(120);default:null"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(100);default:null"`
	Patrons uint32 `json:"patrons" gorm:"type:int unsigned;default:0"`
	Total string `json:"total" gorm:"type:varchar(50);default:null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

}



