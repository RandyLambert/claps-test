package common

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
)

func RegisterType()(gin.HandlerFunc){
	return func(context *gin.Context) {
		gob.Register(github.User{})
		context.Next()
	}
}
