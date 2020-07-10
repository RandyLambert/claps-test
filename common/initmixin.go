package common

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/viper"
)

var MixinClient *mixin.Client
func InitMixin() *mixin.Client {
	s := &mixin.Keystore{
		ClientID:   viper.GetString("client_id"),
		SessionID:  viper.GetString("session_id"),
		PrivateKey: viper.GetString("private_key"),
		PinToken: viper.GetString("pin_token"),
	}

	var err error
	MixinClient,err = mixin.NewFromKeystore(s)
	if err != nil {
		Logger().Error(err.Error())
	}
	return MixinClient
}

func GetMixin() *mixin.Client {
	return MixinClient
}