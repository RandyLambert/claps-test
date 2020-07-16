package controller

import (
	"claps-test/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

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
	user := session.Get("user")
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
	//如果因为超时出错,重新请求

	if err != nil {
		log.Errorf("Users.ListEmails returned error: %v", err)
		ctx.JSON(http.StatusBadRequest,err)
	}

	//根据userId获取所有project信息,Total和Patrons字段添加
	projects,err := service.GetProjectByUserId(1)
	if err != nil {
		log.Errorf("Users.ProjectByUserId returned error: %v", err)
		ctx.JSON(http.StatusBadRequest,err)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"emails": emails,
		"projects": projects,
	})
}

func UserAssets(ctx *gin.Context){

}

func UserTransactions(ctx *gin.Context){

}





