package middleware

import (
	"claps-test/util"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	MixinId string `json:"mixin_id"`
	GithubId string `json:"github_id"`
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
func GenToken(mixin_id ,github_id string) (string, error) {

	c := MyClaims{
		mixin_id,
		github_id,
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
功能:判断请求的Token情况
 */
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		log.Debug("authHeader = ",authHeader)

		//无Token,生成Token返回
		if authHeader == "" {
			log.Debug("无Token")
			token,err := GenToken("","")
			if err != nil{
				c.AbortWithStatusJSON(http.StatusOK,gin.H{
					"message":"generate token error.",
				})
			}

			c.JSON(http.StatusOK,gin.H{
				TOKEN:token,
			})
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
		_, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid Token",
			})
			c.Abort()
			return
		}

		//取出mixin_id和gihub_id
		c.Set(TOKEN,parts[1])

		c.Next()
	}
}


