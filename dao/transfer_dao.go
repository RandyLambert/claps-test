package dao

import (
	"claps-test/model"
	"claps-test/util"
	log "github.com/sirupsen/logrus"
)

func InsertTransfer(transfer *model.Transfer) (err error) {
	err = util.DB.Create(transfer).Error
	return
}

func UpdateTransferTraceId(transferMap *map[string]interface{}, traceId string) (err error) {
	err = util.DB.Debug().Model(model.Transfer{}).Where("trace_id=?", traceId).Updates(*transferMap).Error
	return
}

func GetTransferByUserIdAndAssetId(mixinId string, assetId string) (transfers *[]model.Transfer, err error) {
	transfers = &[]model.Transfer{}
	err = util.DB.Debug().Where("mixin_id = ? AND asset_id = ?", mixinId, assetId).Find(transfers).Error
	log.Debug("dao transfers = ", transfers)
	return
}

//status only '0' or '1' or '2'
func ListTransfersByStatus(status string) (transfer *[]model.Transfer, err error) {
	transfer = &[]model.Transfer{}
	err = util.DB.Where("status=?", status).Find(transfer).Error
	return
}

func UpdateTransferStatusByUserIdAndAssetId(mixinId string, assetId string, status string) (err error) {
	err = util.DB.Debug().Table("transfer").Where("mixin_id = ? AND asset_id= ?", mixinId, assetId).Update("status", status).Error
	return
}

func CountUnfinishedTransfer(mixinId string, assetId string) (count int, err error) {
	err = util.DB.Debug().Table("transfer").Where("mixin_id = ? AND asset_id = ? AND status = ?", mixinId, assetId, model.UNFINISHED).Count(&count).Error
	log.Debug(count)
	return
}
