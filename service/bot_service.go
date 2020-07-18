package service

import (
	"context"
	"github.com/fox-one/mixin-sdk-go"
)

func GetBotAssetByIdAndAssetId(botId string,assetId string)(asset *mixin.Asset,err error){

	client,err := CreateMixinClient(botId)
	if err != nil {
		return
	}
	asset,err = client.ReadAsset(context.Background(),assetId)

	return
}

