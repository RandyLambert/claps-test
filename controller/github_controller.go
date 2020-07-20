package controller

import (
	"claps-test/model"
	"claps-test/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)


func Oauth(ctx *gin.Context){
	log.Debug("开始处理Oauth授权")

	//获取code,path和state
	session := sessions.Default(ctx)
	code := ctx.Query("code")
	path := ctx.Query("path")
	state := ctx.Query("state")


	log.Debug("获得的code,path和state",code,path,state)

	uid,ok := session.Get("uid").(string)
	//不存在state
	ok2 := If(state!="",false,true).(bool)
	if (ok && uid != state ) || ok2 {
		session.Set("user",nil)
		session.Set("githubToken",nil)
		ctx.JSON(http.StatusBadRequest,"invalid oauth redirect")
		return
	}

	//获取token
	var oauthTokenUrl = service.GetOauthToken(code)
	//处理请求的URL,获得Token指针
	token,err := service.GetToken(oauthTokenUrl)
	if err != nil {
		log.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest,"invalid oauth redirect")
		return
	}

	// 通过token，获取用户信息,user是指针类型
	user, err := service.GetUserInfo(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	log.Debugf("\n获得的用户信息是:\n", *user)

	//存储session
	session.Set("user",*user)
	session.Set("githubToken",token.AccessToken)
	session.Save()

	tmp := session.Get("user")
	log.Debug("刚刚存储的session是",tmp)

	//尝试获取数据库中该user信息
	u := model.User{}
	u.Id = *user.ID
	u.Name = *user.Login
	if user.AvatarURL != nil{
		u.AvatarUrl = *user.AvatarURL
	}
	if user.Name != nil{
		u.DisplayName = *user.Name
	}
	if user.Email != nil{
		u.Email = *user.Email
	}

	err = service.InsertOrUpdateUser(&u)
	if err != nil {
		log.Error(err.Error())
	}

	//重定向到http://localhost:3000/profile
	newpath := "http://localhost:3000"+path
	log.Debug("重定向",newpath)
	ctx.Redirect(http.StatusMovedPermanently, newpath)
}
