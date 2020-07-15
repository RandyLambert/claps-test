package util

import (
	"encoding/gob"
	"github.com/google/go-github/v32/github"
)

func RegisterType()(){
	gob.Register(github.User{})
}
