package dao

import (
	"claps-test/model"
	"github.com/jinzhu/gorm"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.Transaction{}).Error; err != nil {
			return err
		}
		return nil
	})
}


func InsertTransaction(transaction *model.Transaction) (err error) {
	err = db.Create(transaction).Error
	return
}

//获取捐赠记录:assetId货币种类,projectName
func ListTransactionsByProjectNameAndAssetId(name string, assetId string) (transactions *[]model.Transaction, err error) {

	transactions = &[]model.Transaction{}
	err = db.Debug().Where("asset_id=? AND project_id=?", assetId,
		db.Debug().Table("project").Select("project.id").Where("project.name=?", name).SubQuery()).Find(transactions).Error
	//err = util.DB.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(transactions).Error
	return
}

//根据项目Id获取项目的所有捐赠
func CountPatronByProjectIdAndSender(projectId uint32, sender string) (count uint32, err error) {
	//todo 可能需要更改
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	err = db.Debug().Table("transaction").Where("project_id=? AND sender=?", projectId, sender).Count(&count).Error
	return
}
