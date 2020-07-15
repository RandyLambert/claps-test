package model

import "time"

type Project struct {
	Id uint32 `json:"id" gorm:"type:int unsigned;primary_key;not null;unique_index:id_UNIQUE"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique_index:name_UNIQUE"`
    DisplayName string `json:"display_name" gorm:"type:varchar(50);default:null"`
	Description string `json:"description" gorm:"type:varchar(120);default:null"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(100);default:null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	//添加两个字段
	/*
	Patrons int64 `json:"patrons"`
	Total float64 `json:"total"`
	 */
}


//project_pro信息
type Project_pro struct {
	Project
	Patrons int64 `json:"patrons"`
	Total float64 `json:"total"`
}


