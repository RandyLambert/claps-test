package controllers

import (
	"claps-test/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func Hello(ctx *gin.Context) {
	//http
	//OauthCallBack := viper.GetString("github.OauthCallBack")
	ctx.HTML(http.StatusOK,"hello.html",gin.H{ //模板渲染
		"ClientID":viper.GetString("github.ClientID"),
		"OauthCallBack":viper.GetString("github.OauthCallBack"),
	})

}

func AuthInfo(ctx *gin.Context){
	//用户信息
}

func Oauth(ctx *gin.Context){
	//获取code
	code := ctx.Query("code")
	//获取token
	var oauthTokenUrl = GetOauthToken(code)
	fmt.Println(oauthTokenUrl)
	token,err := GetToken(oauthTokenUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 通过token，获取用户信息
	var userInfo *models.User
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}
	ctx.JSON(http.StatusOK,*userInfo)

}

func GetOauthToken(code string) string{
	str:= fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		viper.GetString("github.ClientID"),viper.GetString("github.ClientSecret"), code,
		)
	//fmt.Println(str)
	return str
}

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

	if err = json.NewDecoder(res.Body).Decode(&token); err!= nil{
		return nil,err
	}
	return &token,nil
}

func GetUserInfo(token *Token)(*models.User,error){

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
	var userInfo models.User
	body,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body,&userInfo)

	fmt.Printf("+%v",userInfo)
	return &userInfo, nil
}


