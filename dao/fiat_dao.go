package dao

import (
	"claps-test/model"
	"github.com/jinzhu/gorm"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.Fiat{}).Error; err != nil {
			return err
		}
		return nil
	})
}


func UpdateFiat(fiat *model.Fiat) (err error) {
	err = db.Debug().Save(fiat).Error
	return
}

func ListFiatsAllByDB()(fiats *[]model.Fiat,err error){
	fiats = &[]model.Fiat{}
	err = db.Find(fiats).Error
	return
}

func GetFiatByCode(code string)(fiat *model.Fiat,err error) {
	fiat = &model.Fiat{}
	err = db.Select("rate").Where("code = ? ", code).Find(&fiat).Error
	return
}