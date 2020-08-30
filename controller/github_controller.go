package controller

import (
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type oauth struct {
	Code string `json:"code" form:"code"`
	Path string `json:"path" form:"path"`
	State string `json:"state" form:"state"`
}

/*
功能:用code换取Token
说明:此时未完成授权也就没有获得github信息,github_id为空
 */
func Oauth(ctx *gin.Context) {
	var (
		err *util.Err
		oauth_ oauth
		//randomUid = ""
	)

	resp := make(map[string]interface{})

	if err1 := ctx.ShouldBindQuery(&oauth_);err1 != nil ||
		oauth_.Code =="" || oauth_.State == "" || oauth_.Path == ""{
		err1 := util.NewErr(errors.New("invalid query param"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err1, resp)
		return
	}
	log.Debug("code = ",oauth_.Code)
	log.Debug("path = ",oauth_.Path)
	//state设计做为redis的key
	log.Debug("state = ",oauth_.State)

	mcache := &util.MCache{}
	err1 := util.Rdb.Get(oauth_.State,mcache)

	//if err1 == persistence.ErrCacheMiss{
	//验证state
	if err1 != nil{
			err = util.NewErr(err1,util.ErrBadRequest, "")
			util.HandleResponse(ctx, err, resp)
			return
	}


	//获取token
	var oauthTokenUrl = service.GetOauthToken(oauth_.Code)
	//处理请求的URL,获得Token指针
	token2,err := service.GetToken(oauthTokenUrl)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	// 通过token，获取用户信息
	user, err := service.GetUserInfo(token2)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	log.Debug("user = ",*user)

	//redis存储user信息
	mcache.Github = *user
	err1 = util.Rdb.Set(oauth_.Code,mcache,-1)
	if err1 != nil{
		err = util.NewErr(errors.New("cache error"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	tmp := &util.MCache{}
	util.Rdb.Get(oauth_.Code,tmp)
	log.Debug("刚刚存储的user = ",*tmp)

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
	newpath := "http://localhost:3000" + oauth_.Path
	log.Debug("重定向", newpath)
	//ctx.Redirect(http.StatusMovedPermanently, newpath)

}
