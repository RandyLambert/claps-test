package dao

import (
	"claps-test/model"
	"claps-test/util"
)
func InsertAsset(asset *model.Asset)(err error){
	err = util.DB.Create(asset).Error
	return
}

func UpdateAsset(asset *model.Asset)(err error){
	err = util.DB.Save(asset).Error
	return
}

func GetAssetById(assetId string)(asset *model.Asset,err error){
	asset = &model.Asset{}
	err = util.DB.Select("asset_id=?",assetId).Find(asset).Error
	return
}

func ListAssetsAllByDB()(assets *[]model.Asset,err error){
	assets = &[]model.Asset{}
	err = util.DB.Find(assets).Error
	return
}

