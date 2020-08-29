package controller

import (
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Oauth(ctx *gin.Context) {
	var err *util.Err
	resp := make(map[string]interface{})
	log.Debug("开始处理Oauth授权")

	//获取code,path和state
	session := sessions.Default(ctx)
	code := ctx.Query("code")
	path := ctx.Query("path")
	state := ctx.Query("state")

	log.Debug("获得的code,path和state", code, path, state)

	uid, ok := session.Get("uid").(string)
	//不存在state
	ok2 := If(state != "", false, true).(bool)
	if (ok && uid != state) || ok2 {
		session.Set("user", nil)
		session.Set("githubToken", nil)
		err = util.NewErr(errors.New("invalid oauth redirect"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	//获取token
	var oauthTokenUrl = service.GetOauthToken(code)
	//处理请求的URL,获得Token指针
	token, err := service.GetToken(oauthTokenUrl)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	// 通过token，获取用户信息,user是指针类型
	user, err := service.GetUserInfo(token)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	log.Debugf("\n获得的用户信息是:\n", *user)

	//存储session
	session.Set("user", *user)
	session.Set("githubToken", token.AccessToken)
	err1 := session.Save()
	if err1 != nil {
		err = util.NewErr(err1, util.ErrInternalServer, "session保存出错")
		util.HandleResponse(ctx, err, resp)
		return
	}

	tmp := session.Get("user")
	log.Debug("刚刚存储的session是", tmp)

	//尝试获取数据库中该user信息
	u := model.User{}
	u.Id = *user.ID
	u.Name = *user.Login
	if user.AvatarURL != nil {
		u.AvatarUrl = *user.AvatarURL
	}
	if user.Name != nil {
		u.DisplayName = *user.Name
	}
	if user.Email != nil {
		u.Email = *user.Email
	}

	err = service.InsertOrUpdateUser(&u)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	//重定向到http://localhost:3000/profile
	newpath := "http://localhost:3000" + path
	log.Debug("重定向", newpath)
	ctx.Redirect(http.StatusMovedPermanently, newpath)
}
