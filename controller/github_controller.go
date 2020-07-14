package controller

import (
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
	log.Debug("ok2",ok2)
	log.Debug("ok",ok)
	log.Debug("uid",uid)
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
	session.Set("user",*user)
	session.Set("githubToken",token.AccessToken)
	session.Save()

	tmp := session.Get("user")
	log.Debug("刚刚存储的session是",tmp)

	//db := util.GetDB()

	//从数据库中读取user信息

	//没有则插入数据库


	/*
	userInDb := 0
	//尝试获取该user
	db.Debug().Model(model.User{}).Select("id=?",uint32(user["id"].(float64))).Count(&userInDb)
	//没有则插入数据库中
	if userInDb != 1 {
		userInfo := model.User{
			Id:          uint32(user["id"].(float64)),
			Name:        user["login"].(string),
			DisplayName: user["name"].(string),
			Email:       user["email"].(string),
			AvatarUrl:   user["avatar_url"].(string),
		}
		db.Debug().Create(&userInfo)
	}
	//存在则加入session里面
	 */

	//user["envs"] =
	//ctx.JSON(http.StatusOK,user)

	//重定向到http://localhost:3000/profile
	newpath := "http://localhost:3000"+path
	log.Debug("重定向",newpath)
	ctx.Redirect(http.StatusMovedPermanently, newpath)
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

	/*
	userInfoUrl := "https://api.github.com/user"
	req, err := http.NewRequest(http.MethodGet,userInfoUrl,nil)
	if err != nil {
		return nil,err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	var user = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}

	//fmt.Printf("+%v",user)
	return user, nil
	 */
}