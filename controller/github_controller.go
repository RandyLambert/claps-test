package controller

import (
	"claps-test/middleware"
	"claps-test/model"
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-gonic/gin"
)

/**
 * @Description: 用code换取Token,此时没有发放token,成功授权后发放token,token中记录github的userId
 * @param ctx
 */
func Oauth(ctx *gin.Context) {
	type oauth struct {
		Code string `json:"code" form:"code"`
	}
	var (
		err    *util.Err
		oauth_ oauth
	)

	resp := make(map[string]interface{})

	if err1 := ctx.ShouldBindQuery(&oauth_); err1 != nil {
		err := util.NewErr(err1, util.ErrBadRequest, "")
		util.HandleResponse(ctx, err, resp)
		return
	}

	var oauthTokenUrl = service.GetOauthToken(oauth_.Code)
	//处理请求的URL,获得Token指针
	token2, err := service.GetToken(oauthTokenUrl)
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

	//通过token,获取Email信息
	emails, err := service.ListEmailsByToken(token2.AccessToken)
	//如果因为超时出错,重新请求
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	//生成token
	//randomUid := util.RandUp(32)
	randomUid := *user.ID
	jwtToken, err1 := middleware.GenToken(randomUid)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrInternalServer, "gen token error"), nil)
		return
	}


	//向数据库中存储user信息
	u := model.User{}
	u.Id = *user.ID
	u.Name = *user.Login
	if user.AvatarURL != nil {
		u.AvatarUrl = *user.AvatarURL
	}
	if user.Name != nil {
		u.DisplayName = *user.Name
	}
	for _, v := range emails {
		//主email,参与分钱
		if *v.Primary {
			u.Email = *v.Email
			break
		}
	}

	//每次授权后都更新数据库中的信息
	err = service.InsertOrUpdateUser(&u)
	if err != nil {
		util.HandleResponse(ctx, err, resp)
		return
	}

	//token 的uid是github的userId
	resp["token"] = jwtToken
	util.HandleResponse(ctx, nil, resp)
}
