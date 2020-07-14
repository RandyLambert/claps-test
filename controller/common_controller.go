package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"net/http"
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

	//从session中尝试获取用户信息
	session := sessions.Default(ctx)
	var randomUid string = ""

	foxoneToken := session.Get("foxoneToken")
	user := session.Get("user")
	mixinToken := session.Get("mixinToken")

	log.Debug("从session中获取的user",user)
	log.Debug("session中的Token",session.Get("githubToken"))
	log.Debug("从session中获取的mixinToken",mixinToken)
	log.Debug("从session中获取的foxoneToken",foxoneToken)

	if user == nil || mixinToken == nil {
		//没有登录的话随机生成uid
		randomUid = string(RandUp(32))
		//存入session
		session.Set("uid",randomUid)
		session.Save()
	}

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
}

//模拟三目运算符号
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
