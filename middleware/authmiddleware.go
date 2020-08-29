package middleware

import (
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type MyClaims struct {
	MixinId string `json:"mixin_id"`
	GithubId string `json:"github_id"`
	jwt.StandardClaims
}
//jwt的过期时间
const TokenExpireDuration = time.Hour * 2
var MySecret = []byte("claps-dev")

const (
	MIXINID = "mixin_id"
	GITHUBID = "gtihub_id"
	TOKEN = "token"
)

type userInfo struct {
	mixin_id string
	github_id string
}

//判断用户是否登录的中间件
func GithubAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("user")
		if loginuser == nil {
			err := util.NewErr(errors.New("用户没有登录github"), util.ErrUnauthorized, "用户没有登录github")
			util.HandleResponse(ctx, err, nil)
			ctx.Abort()
		} else {
			log.Debug("登录的用户是github", loginuser.(github.User).Name)
			ctx.Next()
		}
	}
}

//判断用户是否登录的中间件
func MixinAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("mixin")
		if loginuser == nil {
			err := util.NewErr(errors.New("用户没有登录mixin"), util.ErrUnauthorized, "用户没有登录mixin")
			util.HandleResponse(ctx, err, nil)
			ctx.Abort()
		} else {
			log.Debug("登录的mixin用户是", loginuser.(string))
			ctx.Next()
		}
	}
}


