package dao

import (
	"claps-test/model"
	"claps-test/util"
	log "github.com/sirupsen/logrus"
)

func InsertTransfer(transfer *model.Transfer)(err error){
	err = util.DB.Create(transfer).Error
	return
}

func UpdateTransfer(transfer *model.Transfer)(err error){
	err = util.DB.Save(transfer).Error
	return

}

func GetTransferByUserIdAndAssetId(userid uint32, assetId string)(transfers *[]model.Transfer,err error) {
	transfers = &[]model.Transfer{}
	err = util.DB.Debug().Where("user_id = ? AND asset_id = ?",userid,assetId).Find(transfers).Error
	log.Debug("dao transfers = ",transfers)
	return
}

//status only '0' or '1'
func ListTransfersByStatus(status rune)(transfer *[]model.Transfer,err error){
	transfer = &[]model.Transfer{}
	err = util.DB.Select("status=?",status).Find(transfer).Error
	return
}
