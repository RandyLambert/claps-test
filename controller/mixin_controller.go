package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-gonic/gin"
)

func Assets(ctx *gin.Context){

	assets,err := service.ListAssetsAllByDB()
	util.HandleResponse(ctx,err,assets)
}
