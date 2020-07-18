package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("user")
		fmt.Println("loginuser:",loginuser)

		if loginuser == nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"没有登录"})
		} else {
			ctx.Next()
		}
	}
}