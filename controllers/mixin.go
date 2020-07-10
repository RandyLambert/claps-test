package controllers

import (
	"claps-test/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Assets(ctx *gin.Context){
	client := common.GetMixin()
	asserts,err := client.ReadAssets(ctx)
	if err != nil {
		log.Error("ReadAssets: ", err.Error())
		ctx.JSON(http.StatusUnauthorized,"Bad Request" + err.Error())
	}
	ctx.JSON(http.StatusOK,&asserts)
}
