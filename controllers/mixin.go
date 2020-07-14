package controllers

import (
	"claps-test/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Assets(ctx *gin.Context){

	assets,err := service.GetAssets()
	if err != nil {
		log.Error("Assets",err.Error())
		ctx.JSON(http.StatusUnauthorized,"Bad Request" + err.Error())
	}

	ctx.JSON(http.StatusOK,assets)
}

//func DoAsset(ctx context.Context, user *sdk.User) string {
	//assetID := "965e5c6e-434c-3fa9-b780-c50f43cd955c"
	//assetID := USDT
	////ReadAsset get asset info, including balance, address info, etc.
	////ReadAsset 获取资产信息，包括余额、地址信息等。
	//asset, err := user.ReadAsset(ctx, assetID)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//printJSON("asset", asset)
	//
	//if asset.AssetID != assetID { //判断是否获取正确
	//	log.Panicf("asset should be %s but get %s\n", assetID, asset.AssetID)
	//}
	//
	//validateAsset(asset)
	//return asset.Destination //返回只资产的充值地址
//}