package model

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段暂时没用到
	Scope       string `json:"scope"`      // 这个字段暂时没用到
}
