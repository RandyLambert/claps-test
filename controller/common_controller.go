package controller

import (
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
 * @Description: 返回环境信息,此时用户没有登录github没有
 * @param ctx
 */
func Environments(ctx *gin.Context) {
	resp := make(map[string]interface{})

	resp["GITHUB_CLIENT_ID"] = util.GithubClinetId
	resp["GITHUB_OAUTH_CALLBACK"] = util.GithubOauthCallback
	resp["MIXIN_CLIENT_ID"] = util.MixinClientId
	resp["MIXIN_OAUTH_CALLBACK"] = util.MixinOauthCallback

	util.HandleResponse(ctx, nil, resp)
}

/**
 * @Description: 认证用户信息,判断github和mixin是否登录绑定,之前有JWTAuthMiddleWare,有jwt说明一定github授权,ctx里设置uid
 * @param ctx
 */
func AuthInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})

	var val interface{}
	var ok bool
	if val, ok = ctx.Get(util.UID); !ok {
		util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), resp)
		return
	}
	uid := val.(int64)

	var mixinAuth bool
	mixinId,err := service.GetMixinIdByUserId(uid)
	if err != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}
	//没有绑定mixin
	if mixinId != ""{
		mixinAuth = true
	}else {
		mixinAuth = false
	}

	user,err := service.GetUserById(uid)
	if err != nil{
		util.HandleResponse(ctx,err,resp)
		return
	}

	resp["user"] = *user
	resp["randomUid"] = uid
	resp["mixinAuth"] = mixinAuth

	util.HandleResponse(ctx, nil, resp)
	return
}

