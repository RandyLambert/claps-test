package service

import (
	"context"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

//从gitub服务器请求获取用户的邮箱信息
func ListEmailsByToken(githubToken string) (err error, emails []*github.UserEmail) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	emails, _, err = client.Users.ListEmails(context.Background(), nil)
	log.Debug("获取的email是",emails)
	return
}
