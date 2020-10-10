package model

import "github.com/jinzhu/gorm"

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&ProjectToMericoGroup{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type ProjectToMericoGroup struct {
	ProjectId     int64  `gorm:"type:bigint;primary_key;not null"`
	MericoGroupId string `gorm:"type:varchar(50);primary_key;not null"`
}
