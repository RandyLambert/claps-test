package controller

import (
	"claps-test/middleware"
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type oauth struct {
	Code string `json:"code"`
	Path string `json:"path"`
	State string `json:"state"`
}

/*
功能:用code换取Token
说明:此时未完成授权也就没有获得github信息,github_id为空
 */
func Oauth(ctx *gin.Context) {
	var (
		err *util.Err
		oauth_ oauth
		randomUid = ""
	)

	resp := make(map[string]interface{})

	if err1 := ctx.ShouldBindQuery(&oauth_);err1 != nil{
		err1 := util.NewErr(errors.New("invalid query param"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err1, resp)
		return
	}
	log.Debug("code = ",oauth_.Code)
	log.Debug("path = ",oauth_.Path)
	log.Debug("state = ",oauth_.State)

	//获取claim
	token1,_ := ctx.Get(middleware.TOKEN)
	jwt_token := token1.(string)
	//claim,_ := middleware.ParseToken(jwt_token)


	err1 := util.Rdb.Get(jwt_token, &randomUid)
	if err1 != nil{
		err = util.NewErr(errors.New("cache error"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	if randomUid != oauth_.State || oauth_.State == ""{
		err = util.NewErr(errors.New("invalid oauth redirect"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	//获取token
	var oauthTokenUrl = service.GetOauthToken(oauth_.Code)
	//处理请求的URL,获得Token指针
	token2,err1 := service.GetToken(oauthTokenUrl)
	if err1 != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	// 通过token，获取用户信息,user是指针类型
	user, err := service.GetUserInfo(token2)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	log.Debug("user = ",*user)

	//redis存储user信息
	github_id := strconv.FormatInt(*user.ID,10)
	err1 = util.Rdb.Set(github_id,*user,-1)
	if err1 != nil{
		err = util.NewErr(errors.New("cache error"), util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	tmp := &github.User{}
	util.Rdb.Get(github_id,tmp)
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

	//不重定向,给用户发发放新的token
	newToken,err1 := middleware.GenToken("",github_id)
	if err1 != nil{
		log.Error("生成Token出错",err1)
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		middleware.TOKEN: newToken,
	})
	/*
	//重定向到http://localhost:3000/profile
	newpath := "http://localhost:3000" + path
	log.Debug("重定向", newpath)
	ctx.Redirect(http.StatusMovedPermanently, newpath)
	 */


}
