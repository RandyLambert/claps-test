package service

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
	"strings"
)

func InsertOrUpdateUser(user *model.User) (err error){
	err = dao.InsertOrUpdateUser(user)
	return
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
func GetToken(url string)(*model.Token,error){

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
	var token model.Token

	//将返回的信息解析到Token
	if err = json.NewDecoder(res.Body).Decode(&token); err!= nil{
		log.Error(err.Error())
		return nil,err
	}
	log.Debug("获得胡Token是",token)
	return &token,nil
}

//用获得的Token获得UserInfo,返回User指针
func GetUserInfo(token *model.Token)(user *github.User,err error){

	log.Debug(token)
	log.Debug("GitHub Token: ",token.AccessToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, _, err = client.Users.Get(ctx, "")


	if err != nil {
		log.Error("n",err)
		return
	}

	return
}

//获取仓库的star数目,如果出错err信息不为空
func GetRepositoryStars(c *gin.Context,slug string)(starCount int,err error) {
	session := sessions.Default(c)
	githubToken := session.Get("githubToken")

	if githubToken == nil{
		log.Error("获取star没有Token")
		return
	}
	log.Debug("获取star数量",githubToken)
	log.Debug("传入的slug是",slug)
	//把slug分成owner和repo
	str := strings.Split(slug,"/")
	log.Debug("owner是",str[0])
	log.Debug("repo是",str[1])

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken.(string)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	var repo *github.Repository
	repo,_,err = client.Repositories.Get(ctx,str[0],str[1])
	starCount =  *repo.StargazersCount
	return
}

