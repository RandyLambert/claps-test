package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段暂时没用到
	Scope       string `json:"scope"`      // 这个字段暂时没用到
}

func Hello(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"hello.html",gin.H{ //模板渲染
		"ClientID":viper.GetString("GITHUB_CLIENT_ID"),
		"OauthCallBack":viper.GetString("GITHUB_OAUTH_CALLBACK"),
	})
}
func UserProfile(ctx *gin.Context){

	ctx.JSON(http.StatusOK,gin.H{
		"emails": 111,
		"projects": 222,
	})
	//session := sessions.Default(ctx)
	//gitHubToken := session.Get("gitHubToken")
	////伪码,具体怎么获得userid还不清楚
	//userId := session.Get("user")
	////获取emails数据

}

func UserAssets(ctx *gin.Context){

}

func UserTransactions(ctx *gin.Context){


}





