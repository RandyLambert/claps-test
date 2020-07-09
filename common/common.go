package common

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

func AuthInfo(ctx *gin.Context){
	//用户信息
	session := sessions.Default(ctx)
	var randomUid string = ""
	user := session.Get("user")
	mixinToken := session.Get("mixinToken")
	if user == nil || mixinToken == nil {
		randomUid = string(RandUp(32))
	}

	ctx.JSON(http.StatusOK,gin.H{"user":user,
		"randomUid":randomUid,
		"mixinAuth":false,
		"foxoneAuth":false,
		"env":gin.H{
			"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
			"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
			"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
		}})
}

