package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MixinAssets(ctx *gin.Context){

	assets,err := service.ListAssetsAllByDB()
	util.HandleResponse(ctx,err,assets)
}

//mixin oauth授权认证
func MixinOauth(ctx *gin.Context)  {

	ctx.JSON(http.StatusOK,gin.H{
		"name":"sky",
		"age":"22",
	})

}
