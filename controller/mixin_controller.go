package controller

import (
	"claps-test/service"
	"claps-test/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func MixinAssets(ctx *gin.Context){

	assets,err := service.ListAssetsAllByDB()
	util.HandleResponse(ctx,err,assets)
}

//mixin oauth授权认证
func MixinOauth(ctx *gin.Context)  {
	code := ctx.Query("code")
	state := ctx.Query("state")
	session := sessions.Default(ctx)

	if code == "" || state == ""{
		err := util.NewErr(nil,util.ErrBadRequest,"mixin oauth认证没有参数")
		util.HandleResponse(ctx,err,nil)
		return
	}

	//判断state和randomUid是否一致
	log.Debug(state)

	//用state换取令牌
	client,err := service.GetMixinAuthorizedClient(ctx,code)
	if err != nil{
		util.HandleResponse(ctx,err,nil)
		return
	}

	//获取mixin用户信息,存入session
	user,err2 := service.GetMixinUserInfo(ctx,client)
	if err2 != nil{
		util.HandleResponse(ctx,err2,nil)
		return
	}
	session.Set("mixin",client)



	log.Println("user", user.UserID)

	//github一定是登录,绑定mixin和github


	//重定位
	//ctx.Redirect(http.StatusOK,"www.baidu.com")

	ctx.JSON(http.StatusOK,gin.H{
		"data":user,
	})

}
