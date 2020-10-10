package model

import "github.com/jinzhu/gorm"

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Member{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Member struct {
	ProjectId int64 `gorm:"type:bigint;not null;primary_key"`
	UserId    int64 `gorm:"type:bigint;not null;primary_key"`
}
