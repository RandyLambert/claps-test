package service

import (
	"claps-test/util"
	"context"
	"errors"
	"github.com/fox-one/mixin-sdk-go"
	log "github.com/sirupsen/logrus"
)

func GetAssets()(assets []*mixin.Asset,err error){

	assets,err = util.MixinClient.ReadAssets(context.Background())
	return
}

func GetAsset(assetID string) (asset *mixin.Asset,err error){
	asset,err =  util.MixinClient.ReadAsset(context.Background(),assetID)
	if asset != nil{
		if assetID != asset.AssetID {
			log.Error("asset should be %s but get %s\n", assetID, asset.AssetID)
			return nil,errors.New("GetAsset error")
		}
	}
	return
}

