package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func InsertTransfer(transfer *model.Transfer)(err error){
	err = util.DB.Create(transfer).Error
	return
}