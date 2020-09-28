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

//获取捐赠记录:通过projectName
func ListTransactionsByProjectId(projectId string) (transactions *[]model.Transaction, err error) {

	transactions = &[]model.Transaction{}
	err = db.Debug().Where("project_id=?", projectId).Order("created_at desc").Limit(256).Find(transactions).Error

	return
}

//根据项目Id获取项目的所有捐赠
func CountPatronByProjectIdAndSender(projectId string, sender string) (count uint32, err error) {
	//todo 可能需要更改
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	err = db.Debug().Table("transaction").Where("project_id=? AND sender=?", projectId, sender).Count(&count).Error
	return
}
