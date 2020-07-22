package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-gonic/gin"
)

func MixinAssets(ctx *gin.Context){

	assets,err := service.ListAssetsAllByDB()
	util.HandleResponse(ctx,err,assets)
}
