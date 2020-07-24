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

//status only '0' or '1' or '2'
func ListTransfersByStatus(status string)(transfer *[]model.Transfer,err error){
	transfer = &[]model.Transfer{}
	err = util.DB.Where("status=?",status).Find(transfer).Error
	return
}

func UpdateTransferSatusByUserIdAndAssetId(userId uint32, assetId string,status string)(err error)  {
	err = util.DB.Debug().Table("transfer").Where("user_id = ? AND asset_id= ?",userId,assetId).Update("status",status).Error
	return
}

func CountUnfinishedTransfer(userId uint32, assetId string) (count int,err error) {
	err = util.DB.Debug().Table("transfer").Where("user_id = ? AND asset_id = ? AND status = ?",userId,assetId,model.UNFINISHED).Count(&count).Error
	log.Debug(count)
	return
}