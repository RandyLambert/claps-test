package util

import (
	"encoding/json"
	"github.com/google/go-github/v32/github"
)

func UserToJson(user *github.User)(userJson string,err error)  {
	jsonBytes,err := json.Marshal(*user)
	userJson = string(jsonBytes)
	return
}

func JsonToUser(userJson string)(user *github.User,err error)  {
	user = &github.User{}
	err = json.Unmarshal([]byte(userJson), &user)
	return
}