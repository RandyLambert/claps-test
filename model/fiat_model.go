package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Fiat{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Fiat struct {
	Code string          `json:"code,omitempty" gorm:"type:varchar(25);primary_key;not null"`
	Rate decimal.Decimal `json:"rate,omitempty" gorm:"type:varchar(128);not null"`
}

var FIAT *Fiat

func (fiat *Fiat) UpdateFiat(fiatData *Fiat) (err error) {
	err = db.Debug().Save(fiatData).Error
	return
}

func (fiat *Fiat) ListFiatsAllByDB() (fiats *[]Fiat, err error) {
	fiats = &[]Fiat{}
	err = db.Find(fiats).Error
	return
}

func (fiat *Fiat) GetFiatByCode(code string) (fiatData *Fiat, err error) {
	fiatData = &Fiat{}
	err = db.Select("rate").Where("code = ? ", code).Find(&fiat).Error
	return
}
