package controller

import (
	"claps-test/dao"
	"claps-test/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"net/http"
)

func Oauth(ctx *gin.Context){
	log.Debug("开始处理Oauth授权")

	//获取code,path和state
	session := sessions.Default(ctx)
	code := ctx.Query("code")
	path := ctx.Query("path")
	state := ctx.Query("state")


	log.Debug("获得的code,path和state",code,path,state)

	uid,ok := session.Get("uid").(string)
	//不存在state
	ok2 := If(state!="",false,true).(bool)
	if (ok && uid != state ) || ok2 {
		session.Set("user",nil)
		session.Set("githubToken",nil)
		ctx.JSON(http.StatusBadRequest,"invalid oauth redirect")
		return
	}

	//获取token
	var oauthTokenUrl = GetOauthToken(code)
	//处理请求的URL,获得Token指针
	token,err := GetToken(oauthTokenUrl)
	if err != nil {
		log.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest,"invalid oauth redirect")
		return
	}

	// 通过token，获取用户信息,user是指针类型
	user, err := GetUserInfo(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	log.Debug("\n获得的用户信息是:\n", *user)

	//存储session
	session.Set("gxk","gaoxingkun")
	session.Save()
	session.Set("user",*user)
	session.Set("githubToken",token.AccessToken)
	session.Save()

	tmp := session.Get("user")
	log.Debug("刚刚存储的session是",tmp)

	//尝试获取数据库中该user信息
	u := model.User{}
	u.Id = *user.ID
	u.Name = *user.Login
	if user.AvatarURL != nil{
		u.AvatarUrl = *user.AvatarURL
	}
	if user.Name != nil{
		u.DisplayName = *user.Name
	}
	if user.Email != nil{
		u.Email = *user.Email
	}

	CreateOrUpdateUser(&u)


	//重定向到http://localhost:3000/profile
	newpath := "http://localhost:3000"+path
	log.Debug("重定向",newpath)
	ctx.Redirect(http.StatusMovedPermanently, newpath)
}

func CreateOrUpdateUser(user *model.User) {
	dao.CreateOrUpdateUser(user)
}

func GetUserById(user *model.User,Id int64){
	dao.SelectUserById(user,Id)
}

/*
拼接含有code和clientID和client_secret，成一个URL用来换取Token,返回一个拼接的URL
code 表示github认证服务器返回的code
 */
func GetOauthToken(code string) string{
	str:= fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		viper.GetString("GITHUB_CLIENT_ID"),viper.GetString("GITHUB_CLIENT_SECRET"), code,
	)
	//fmt.Println(str)
	return str
}

/*
根据参数URL去请求，然后换取Token,返回Token指针和错误信息
 */
func GetToken(url string)(*Token,error){

	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("accept","application/json")

	//发送请求并获得响应
	var httpClient = http.Client{}

	res,err := httpClient.Do(req);
	if err != nil {
		return nil,err
	}

	//将相应体解析为token,返回
	var token Token

	//将返回的信息解析到Token
	if err = json.NewDecoder(res.Body).Decode(&token); err!= nil{
		log.Error(err.Error())
		return nil,err
	}
	log.Debug("获得胡Token是",token)
	return &token,nil
}

//用获得的Token获得UserInfo,返回User指针
func GetUserInfo(token *Token)(*github.User,error){

	log.Debug(token)
	log.Debug("GitHub Token: ",token.AccessToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, "")


	if err != nil {
		log.Error("n",err)
		return user,err
	}

	return user,err

}