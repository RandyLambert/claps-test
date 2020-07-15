package controller

import (
	"claps-test/model"
	"claps-test/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
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


	//project_pro信息
	type project_pro struct {
		model.Project
		Patrons int64 `json:"patrons"`
		Total float64 `json:"total"`
	}

	//根据userId获取所有project信息,Total和Patrons字段添加



	mp := make(map[string]interface{})

	//给每个项目添加Totol字段和patrons字段
	projects := []project_pro{}
	p := project_pro{}
	p.Name = "claps.dev"
	p.Patrons = 1
	p.Id = 1
	p.DisplayName = "Claps.dev"
	p.AvatarUrl = "http://dmimg.5054399.com/allimg/pkm/pk/13.jpg"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Description = "abc abc asdw qwerasdfzxc "
	p.Total = 0.1234
	projects = append(projects,p )

	/*
	"total":4.6176027,
		"patrons":1,
		"id":1,
		"name":"claps.dev",
		"displayName":"Claps.dev",
		"description":"abc",
		"avatarUrl":"abc",
		"createdAt":"2020-04-06T03:46:08.000Z",
		"updatedAt":"2020-04-06T03:46:08.000Z"
	 */
	ctx.JSON(http.StatusOK,gin.H{
		"emails": emails,
		"projects": projects,
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





