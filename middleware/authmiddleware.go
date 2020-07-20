package middleware

import (
	"claps-test/model"
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//判断用户是否登录的中间件
func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("user")

		if loginuser == nil{
			err := util.NewErr(errors.New("用户没有登录"),util.ErrUnauthorized,"用户没有登录")
			util.HandleResponse(ctx,err,nil)
		} else {
			log.Debug("登录的用户是",loginuser.(model.User).Name)
			ctx.Next()
		}
	}
}