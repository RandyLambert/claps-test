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

//获取捐赠记录:通过projectId
func ListTransactionsByProjectId(projectId string) (transactions *[]model.Transaction, err error) {

	transactions = &[]model.Transaction{}
	err = db.Debug().Where("project_id=?", projectId).Order("created_at desc").Limit(256).Find(transactions).Error

	return
}
