package models

type User struct{
	Id uint32 `gorm:"type:int unsigned;primary_key;unique_index:id_UNIQUE;not null;"`
	Name string `gorm:"type:varchar(50);unique_index:name_UNIQUE;not null"`
	DisplayName string `gorm:"type:varchar(50);default:null"`
	Email string `gorm:"type:varchar(50);unique_index:email_UNIQUE;not null"`
	AvatarUrl string `gorm:"type:varchar(100);default:null"`
}
