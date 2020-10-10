package model

import "github.com/jinzhu/gorm"

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Property{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//用来做重要数据的持久化,保存key-value
type Property struct {
	Key   string `gorm:"type:varchar(50);not null;default:0;primary_key;"`
	Value string `gorm:"type:varchar(50);not null;default:0"`
}

var PROPERTY *Property

func (prop *Property) GetPropertyByKey(Key string) (property *Property, err error) {
	property = &Property{
		Key: Key,
	}
	err = db.First(property).Error
	return
}

func (prop *Property) UpdateProperty(property *Property) (err error) {
	err = db.Save(property).Error
	return
}
