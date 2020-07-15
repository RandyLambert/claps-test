package controller

import (
	"claps-test/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

//获取用户的邮箱和项目
func UserProfile(ctx *gin.Context){

	session := sessions.Default(ctx)
	//从session中获取user和githubToken信息
	user:=  session.Get("user")
	githubToken := session.Get("githubToken")

	if user == nil || githubToken == nil{
		log.Error("session中信息为空")
		//重定向登录
		ctx.Redirect(400,"http://localhost:3000/")
	}

	token := githubToken.(string)
	log.Debugf("\n\nprofile中的user:",user)
	log.Debugf("\n\nprofile中的githubToken:",token)

	//获取email信息
	err,emails := service.GetEmailInfo(token)
	if err != nil {
		log.Errorf("Users.ListEmails returned error: %v", err)
		ctx.JSON(http.StatusBadRequest,err)
	}

	//获取project信息
	var tmp []interface{}
	ctx.JSON(http.StatusOK,gin.H{
		"emails": emails,
		"projects": tmp,
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





