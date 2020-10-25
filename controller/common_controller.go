package controller

import (
	"claps-test/middleware"
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)


/*
功能:返回环境信息
说明:此时用户没有登录github没有
*/
func Environments(ctx *gin.Context) {
	resp := make(map[string]interface{})

	resp["GITHUB_CLIENT_ID"] = util.GithubClinetId
	resp["GITHUB_OAUTH_CALLBACK"] = util.GithubOauthCallback
	resp["MIXIN_CLIENT_ID"] = util.MixinClientId
	resp["MIXIN_OAUTH_CALLBACK"] = util.MixinOauthCallback

	util.HandleResponse(ctx, nil, resp)
}

/*
功能:认证用户信息,判断github和mixin是否登录绑定
说明:之前有JWTAuthmiddleWare,有jwt说明一定github授权,ctx里设置uid
*/
func AuthInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})

	//获取claim
	uid, _ := ctx.Get(util.UID)
	randomUid, _ := uid.(string)
	log.Debug("RandomUid = ", randomUid)

	//从redis取出mcache
	mcache := &util.MCache{}
	err1 := util.Rdb.Get(randomUid, mcache)
	if err1 != nil {
		util.HandleResponse(ctx, util.NewErr(err1, util.ErrInternalServer, "get cacche error"), nil)
		return
	}

	//cache中没有mixin信息
	if !mcache.MixinAuth {
		//更新mixin信息
		mixin_id, err := service.GetMixinIdByUserId(*mcache.Github.ID)
		if err != nil {
			util.HandleResponse(ctx, err, nil)
			return
		}

		if mixin_id != "" {
			//set cache ,next
			mcache.MixinId = mixin_id
			mcache.MixinAuth = true
			err1 = util.Rdb.Replace(randomUid, *mcache, -1)
			if err1 != nil {
				err = util.NewErr(errors.New("cache set error"), util.ErrDataBase, "")
				util.HandleResponse(ctx, err, nil)
				return
			}
		}
	}

	//从redis中取出github信息返回
	resp["user"] = mcache.Github
	resp["randomUid"] = uid
	resp["mixinAuth"] = mcache.MixinAuth

	util.HandleResponse(ctx, nil, resp)
	return
}

/*
功能:再无Token的情况下,返回Uid和Token,并且redis缓存uid-mcache
*/
func noToken(c *gin.Context) (randomUid string) {
	resp := make(map[string]interface{})
	randomUid = util.RandUp(32)

	token, err := middleware.GenToken(randomUid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "generate token error.",
		})
	}
	resp["token"] = token

	mcache := util.MCache{}
	err1 := util.Rdb.Set(randomUid, mcache, -1)
	if err1 != nil {
		util.HandleResponse(c, util.NewErr(err1, util.ErrDataBase, "cache set error"), nil)
		return
	}

	util.HandleResponse(c, nil, resp)
	return
}

/**
 * @Description: 模拟三目运算符号
 * @param condition
 * @param trueVal
 * @param falseVal
 * @return interface{}
 */
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
