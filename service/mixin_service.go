package service

import (
	"claps-test/common"
	"context"
	"errors"
	"github.com/fox-one/mixin-sdk-go"
	log "github.com/sirupsen/logrus"
)

func GetAssets()(assets []*mixin.Asset,err error){

	assets,err = common.MixinClient.ReadAssets(context.Background())
	return
}

func GetAsset(assetID string) (asset *mixin.Asset,err error){
	asset,err =  common.MixinClient.ReadAsset(context.Background(),assetID)
	if asset != nil{
		if assetID != asset.AssetID {
			log.Error("asset should be %s but get %s\n", assetID, asset.AssetID)
			return nil,errors.New("GetAsset error")
		}
	}
	return
}
