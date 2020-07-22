package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func InsertTransfer(transfer *model.Transfer)(err error){
	err = util.DB.Create(transfer).Error
	return
}

func UpdateTransfer(transfer *model.Transfer)(err error){
	err = util.DB.Save(transfer).Error
	return

}

//status only '0' or '1' or '2'
func ListTransfersByStatus(status rune)(transfer *[]model.Transfer,err error){
	transfer = &[]model.Transfer{}
	err = util.DB.Where("status=?",status).Find(transfer).Error
	return
}
