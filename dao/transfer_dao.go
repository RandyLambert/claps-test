package dao

import (
	"claps-test/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.Transfer{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func InsertOrUpdateTransfer(transfer *model.Transfer) (err error) {
	err = db.Debug().Save(transfer).Error
	return
}

func GetTransferByMixinId(mixinId string) (transfers *[]model.Transfer, err error) {
	transfers = &[]model.Transfer{}
	err = db.Debug().Where("mixin_id = ?", mixinId).Find(transfers).Error
	return
}

//status only '0' or '1'
func ListTransfersByStatus(status string) (transfer *[]model.Transfer, err error) {
	transfer = &[]model.Transfer{}
	err = db.Where("status=?", status).Find(transfer).Error
	return
}

func CountUnfinishedTransfer(mixinId string) (count int, err error) {
	err = db.Debug().Table("transfer").Where("status = ? AND mixin_id = ?", model.UNFINISHED, mixinId).Count(&count).Error
	log.Debug(count)
	return
}
