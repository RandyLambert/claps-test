package service

import (
	"claps-test/util"
	"context"
	"github.com/fox-one/mixin-sdk-go"
)

func GetAssetByBotIdAndAssetId(botId string,assetId string)(asset *mixin.Asset,err *util.Err){

	client,err := CreateMixinClient(botId)
	if err != nil {
		return
	}
	asset,err1 := client.ReadAsset(context.Background(),assetId)
	if err1 != nil {
		err = util.NewErr(err,util.ErrThirdParty,"通过botid读取asset信息失败")
	}
	return
}

