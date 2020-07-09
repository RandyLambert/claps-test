package controllers

import (
	"claps-test/dao"
	"claps-test/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Oauth(ctx *gin.Context){
	//获取code
	code := ctx.Query("code")
	//获取token
	var oauthTokenUrl = GetOauthToken(code)
	//fmt.Println(oauthTokenUrl)
	token,err := GetToken(oauthTokenUrl)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest,"invalid oauth redirect")
		return
	}
	// 通过token，获取用户信息

	user, err := GetUserInfo(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	db := dao.GetDB()
	userInDb := 0
	db.Debug().Model(models.User{}).Select("id=?",uint32(user["id"].(float64))).Count(&userInDb)
	if userInDb != 1 {
		userInfo := models.User{
			Id:          uint32(user["id"].(float64)),
			Name:        user["login"].(string),
			DisplayName: user["name"].(string),
			Email:       user["email"].(string),
			AvatarUrl:   user["avatar_url"].(string),
		}
		db.Debug().Create(&userInfo)
	}

	//user["envs"] =

	ctx.JSON(http.StatusOK,user)

}

func GetOauthToken(code string) string{
	str:= fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		viper.GetString("GITHUB_CLIENT_ID"),viper.GetString("GITHUB_CLIENT_SECRET"), code,
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

func GetUserInfo(token *Token)(map[string]interface{},error){

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
}