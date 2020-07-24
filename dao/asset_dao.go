package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func UpdateAsset(asset *model.Asset)(err error){
	err = util.DB.Save(asset).Error
	return
}

func GetAssetById(assetId string)(asset *model.Asset,err error){
	asset = &model.Asset{}
	err = util.DB.Where("asset_id=?",assetId).Find(asset).Error
	return
}

func ListAssetsAllByDB()(assets *[]model.Asset,err error){
	assets = &[]model.Asset{}
	err = util.DB.Find(assets).Error
	return
}

func GetPriceUsdByAssetId(assetId string)(priceUsd *model.AssetIdToPriceUsd,err error){
	priceUsd = &model.AssetIdToPriceUsd{}
	err = util.DB.Debug().Table("asset").Select("asset_id,price_usd").Where("asset_id=?",assetId).Scan(priceUsd).Error
	return
}
