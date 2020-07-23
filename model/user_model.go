package model

type User struct{
	Id uint32`json:"id" gorm:"type:int unsigned;primary_key;not null;"`
	Name string `json:"login" gorm:"type:varchar(50);unique_index:name_UNIQUE;not null"`
	DisplayName string `json:"name" gorm:"type:varchar(50);default:null"`
	Email string `json:"email" gorm:"type:varchar(50);unique_index:email_UNIQUE;not null"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(100);default:null"`
	//UserId string `json:"user_id" gorm:"type:varchar(50);default:null"`
	MixinId string `json:"user_id" gorm:"type:varchar(50);default:null"`
}

type UserMixinId struct {
	Id uint32`json:"id" gorm:"type:int unsigned;primary_key;not null;"`
	MixinId string `json:"user_id" gorm:"type:varchar(50);default:null"`
}