package service

import (
	"claps-test/dao"
	"context"
	"github.com/fox-one/mixin-sdk-go"
)

func GetBotAsset(botId string,assetId string)(asset *mixin.Asset,err error){

	client,err := CreateMixinClient(botId)
	if err != nil {
		return
	}
	asset,err = client.ReadAsset(context.Background(),assetId)

	return
}

func CreateMixinClient(botId string)(client *mixin.Client,err error){
	bot,err := dao.GetBot(botId)
	if err != nil {
		return
	}
	s := &mixin.Keystore{
		ClientID:   bot.Id,
		SessionID:  bot.SessionId,
		PrivateKey: bot.PrivateKey,
		PinToken: bot.PinToken,
	}


	client, err = mixin.NewFromKeystore(s)
	return
}
