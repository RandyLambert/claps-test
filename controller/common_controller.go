package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
)
var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")

func RandUp(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 63
		b[i] = longLetters[arc]
	}
	return b
}

//认证身份信息
func AuthInfo(ctx *gin.Context){
	resp := make(map[string]interface{})
	//从session中尝试获取用户信息
	session := sessions.Default(ctx)
	var randomUid string = ""


	//foxoneToken := session.Get("foxoneToken")
	user := session.Get("user")
	//mixin中存储的是用户mixin的user_id
	mixinToken := session.Get("mixin")

	//如果session中没有mixin的user_id尝试从数据库读取,如果绑定了就不需要用户在绑定mixin了
	if user != nil && mixinToken == nil{
		userId := uint32(*user.(github.User).ID)
		mixinId,err := service.GetMixinIdByUserId(userId)
		if err != nil{
			util.HandleResponse(ctx,err,nil)
			return
		}
		if mixinId != ""{
			session.Set("mixin",mixinId)
			err := session.Save()
			if err != nil{
				util.HandleResponse(ctx,util.NewErr(err,util.ErrInternalServer,"保存session出错"),nil)
				return
			}
			mixinToken = true
		}
	}


	log.Debug("从session中获取的user",user)
	log.Debug("从session中获取的mixinToken",mixinToken)
	//log.Debug("从session中获取的foxoneToken",foxoneToken)

	if user == nil || mixinToken == nil {
		//没有登录的话随机生成uid
		randomUid = string(RandUp(32))
		//存入session
		session.Set("uid",randomUid)
		err1 := session.Save()
		if err1 != nil{
			err := util.NewErr(err1,util.ErrInternalServer,"session保存出错")
			util.HandleResponse(ctx,err,resp)
			return
		}
	}

	resp["user"] = user
	resp["randomUid"] = randomUid
	resp["mixinAuth"] = If(mixinToken != nil,true,false).(bool)
	//resp["foxoneAuth"] = If(foxoneToken!= nil,true,false).(bool)
	resp["envs"] = gin.H{
		"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
		"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
		"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
	}

	util.HandleResponse(ctx,nil,resp)

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
