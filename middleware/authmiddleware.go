package middleware

import (
	"claps-test/util"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	//MixinId string `json:"mixin_id"`
	//GithubId string `json:"github_id"`
	Uid string `json:"uid"`
	jwt.StandardClaims
}


var MySecret = []byte("claps-dev")

const (
	MIXINID = "mixin_id"
	GITHUBID = "gtihub_id"
	TOKEN = "token"
)

type userInfo struct {
	mixin_id string
	github_id string
}

//判断用户是否登录的中间件
func GithubAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("user")
		if loginuser == nil {
			err := util.NewErr(errors.New("用户没有登录github"), util.ErrUnauthorized, "用户没有登录github")
			util.HandleResponse(ctx, err, nil)
			ctx.Abort()
		} else {
			log.Debug("登录的用户是github", loginuser.(github.User).Name)
			ctx.Next()
		}
	}
}

//判断用户是否登录的中间件
func MixinAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取session
		session := sessions.Default(ctx)
		loginuser := session.Get("mixin")
		if loginuser == nil {
			err := util.NewErr(errors.New("用户没有登录mixin"), util.ErrUnauthorized, "用户没有登录mixin")
			util.HandleResponse(ctx, err, nil)
			ctx.Abort()
		} else {
			log.Debug("登录的mixin用户是", loginuser.(string))
			ctx.Next()
		}
	}
}

/*
功能:生成Tokenm
参数:mixin的userID和github的Id
 */
func GenToken(uid string) (string, error) {

	c := MyClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(util.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "sky",                               // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)

}

/*
功能:解析jwt为Myclaim
参数:jwt字符号
 */
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/*
功能:再无Token的情况下,返回Uid和Token,并且redis缓存uid-mcache
 */
func NoToken(c *gin.Context)(randomUid string)  {
	resp := make(map[string]interface{})
	randomUid = util.RandUp(32)


	token,err := GenToken(randomUid)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusOK,gin.H{
			"message":"generate token error.",
		})
	}
	resp["user"] = nil
	resp["randomUid"] = randomUid
	resp["mixinAuth"] = false
	resp["envs"] = gin.H{
		"GITHUB_CLIENT_ID":      viper.GetString("GITHUB_CLIENT_ID"),
		"GITHUB_OAUTH_CALLBACK": viper.GetString("GITHUB_OAUTH_CALLBACK"),
		"MIXIN_CLIENT_ID":       viper.GetString("MIXIN_CLIENT_ID"),
	}
	resp["token"] = token

	mcache := util.MCache{}

	err1 := util.Rdb.Set(randomUid,mcache,-1)
	if err1 != nil{
		util.HandleResponse(c,util.NewErr(err1,util.ErrDataBase,"cache set error"),nil)
		return
	}

	util.HandleResponse(c,nil,resp)
	return
}

/*
功能:判断请求的Token情况
 */
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		log.Debug("authHeader = ",authHeader)

		var randomUid string
		//无Token,生成Token返回,生成Uid
		if authHeader == "" {
			log.Debug("No Token")
			randomUid = NoToken(c)
			fmt.Println("randomUid = ",randomUid)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid Token",
			})
			c.Abort()
			return
		}

		//set Key
		c.Set(util.UID,claim.Uid)
		c.Next()
	}
}


