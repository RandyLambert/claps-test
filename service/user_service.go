package service

import (
	"claps-test/util"
	"context"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

//从gitub服务器请求获取用户的邮箱信息
func ListEmailsByToken(githubToken string) (emails []*github.UserEmail,err *util.Err) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	emails, _, err2 := client.Users.ListEmails(context.Background(), nil)

	if err2 != nil{
		err = util.NewErr(err2,util.ErrThirdParty,"从github获取Email错误")
	} else{
		err = util.OK
		log.Debug("获取的email是",emails)
	}

	return
}
