package model

type User struct{
	Id int64`json:"id" gorm:"type:int unsigned;primary_key;unique_index:id_UNIQUE;not null;"`
	Name string `json:"login" gorm:"type:varchar(50);unique_index:name_UNIQUE;not null"`
	DisplayName string `json:"name" gorm:"type:varchar(50);default:null"`
	Email string `json:"email" gorm:"type:varchar(50);unique_index:email_UNIQUE;not null"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(100);default:null"`
}
