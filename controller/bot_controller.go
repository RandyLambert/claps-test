package controller

import (
	"claps-test/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Bot(ctx *gin.Context){

	asset,err := service.GetBotAssetByIdAndAssetId(ctx.Param("botId"),ctx.Param("assetId"))

	if err != nil {
		log.Error("Bot",err.Error())
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
	}else {
		ctx.JSON(http.StatusOK,asset)
	}

}
