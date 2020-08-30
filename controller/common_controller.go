package controller

import (
	"claps-test/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	RANDOMUID = "randomUid"
)


/*
功能:返回前端randomUid
说明:调用此函数时候用户没有登录github,生成randomUid,并存储redis
 */
func LoginGithub(ctx *gin.Context)  {
	resp := make(map[string]interface{})
	uid,_ := ctx.Get(util.UID)
	randomUid,_ := uid.(string)

	resp["user"] = nil
	resp["randomUid"] = randomUid
	resp["mixinAuth"] = false
	resp["envs"] = gin.H{
		"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
		"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
		"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
	}
	util.HandleResponse(ctx, nil, resp)
	return
}
/*
功能:认证用户信息,判断github和mixin是否登录绑定
说明:之前有JWTAuthmiddleWare,ctx里设置uid
 */
func AuthInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})

	/*
	//从session中尝试获取用户信息
	session := sessions.Default(ctx)
	var randomUid string = ""

	//foxoneToken := session.Get("foxoneToken")
	user := session.Get("user")
	//mixin中存储的是用户mixin的user_id
	mixinToken := session.Get("mixin")
	 */
	/*
	var (
		randomUid = ""
		user  *github.User = nil
	)

	 */

	//获取claim
	uid,_ := ctx.Get(util.UID)
	randomUid,_ := uid.(string)

	log.Debug("RandomUid = ",randomUid)

	//从redis取出mcache
	mcache := &util.MCache{}
	err1 := util.Rdb.Get(randomUid,mcache)
	if err1 != nil{
		log.Error("Get cache error:",err1)
		return
	}

	//未登录github
	if mcache.GithubAuth == false{
		log.Debug("github未登录")
		LoginGithub(ctx)
		return
	}


	log.Debug("github已登录")
	//从redis中取出github信息返回
	resp["user"] = mcache.Github
	resp["randomUid"] = randomUid
	resp["mixinAuth"] = mcache.MixinAuth
	resp["envs"] = gin.H{
		"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
		"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
		"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
	}

	util.HandleResponse(ctx, nil, resp)
	return


	/*
	//用户已经登录github
	if github_id != ""{
		userId := *user.(github.User).ID
		//通过github_id获取mixin_id
		mixinId, err := service.GetMixinIdByUserId(userId)
		if err != nil {
			util.HandleResponse(ctx, err, nil)
			return
		}

		if mixinId != "" {
			session.Set("mixin", mixinId)
			err := session.Save()
			if err != nil {
				util.HandleResponse(ctx, util.NewErr(err, util.ErrInternalServer, "保存session出错"), nil)
				return
			}
			mixinToken = true
		}
	}
	 */

	/*
	//如果session中没有mixin的user_id尝试从数据库读取,如果绑定了就不需要用户在绑定mixin了
	if user != nil && mixinToken == nil {
		userId := *user.(github.User).ID
		mixinId, err := service.GetMixinIdByUserId(userId)
		if err != nil {
			util.HandleResponse(ctx, err, nil)
			return
		}
		if mixinId != "" {
			session.Set("mixin", mixinId)
			err := session.Save()
			if err != nil {
				util.HandleResponse(ctx, util.NewErr(err, util.ErrInternalServer, "保存session出错"), nil)
				return
			}
			mixinToken = true
		}
	}
	 */

	//if user == nil || mixinToken == nil {
	//未登录github或者未绑定mixin
	//if mixin_id == ""|| github_id== ""{
	//		//没有登录的话随机生成uid
	//		randomUid = string(RandUp(32))
	//		//存入redis
	//		util.Rdb.Set(RANDOMUID,randomUid,middleware.TokenExpireDuration)
	//	}
	//
	//resp["user"] = user
	//resp["randomUid"] = randomUid
	////resp["mixinAuth"] = If(mixinToken != nil, true, false).(bool)
	//resp["envs"] = gin.H{
	//	"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
	//	"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
	//	"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
	//}

	util.HandleResponse(ctx, nil, resp)

	/*
		ctx.JSON(http.StatusOK,gin.H{
			"user":user,
			"randomUid":randomUid,
			"mixinAuth": If(mixinToken != nil,true,false).(bool),
			"foxoneAuth": If(foxoneToken!= nil,true,false).(bool),
			"envs":gin.H{
				"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
				"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
				"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
			}})
	*/
}

//模拟三目运算符号
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
