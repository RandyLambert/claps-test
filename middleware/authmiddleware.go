package middleware

import (
	"claps-test/service"
	"claps-test/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	MIXINID  = "mixin_id"
	GITHUBID = "gtihub_id"
	TOKEN    = "token"
)

type userInfo struct {
	mixin_id  string
	github_id string
}

/**
 * @Description: 判断github是否已经授权,经过了JWT中间件,一定有cache key
 * @return gin.HandlerFunc
 */
func GithubAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var val interface{}
		var ok bool
		if val, ok = ctx.Get(util.UID); !ok {
			util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), nil)
			return
		}
		uid := val.(string)

		mcache := &util.MCache{}
		err1 := util.Rdb.Get(uid, mcache)
		if err1 != nil {
			util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), nil)
			return
		}

		//github未登录
		if !mcache.GithubAuth {
			util.HandleResponse(ctx, util.NewErr(err1, util.ErrUnauthorized, "github unauthorized"), nil)
			return
		}
		ctx.Next()
	}
}
/**
 * @Description: 检查是够绑定mixin,github一定是登录了,从数据库中查询问是否绑定mixin,绑定则更新缓存
 * @return gin.HandlerFunc
 */
func MixinAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var val interface{}
		var ok bool
		if val, ok = ctx.Get(util.UID); !ok {
			util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrDataBase, "ctx get uid error"), nil)
			return
		}
		uid := val.(string)

		mcache := &util.MCache{}
		err1 := util.Rdb.Get(uid, mcache)
		if err1 != nil {
			util.HandleResponse(ctx, util.NewErr(err1, util.ErrDataBase, "cache get error"), nil)
			return
		}

		if mcache.MixinAuth {
			ctx.Next()
		}

		//从数据库查询mixin_id
		mixinId, err := service.GetMixinIdByUserId(*mcache.Github.ID)
		if err != nil {
			util.HandleResponse(ctx, err, nil)
			ctx.Abort()
		}

		if mixinId == "" {
			util.HandleResponse(ctx, util.NewErr(err1, util.ErrUnauthorized, "mixin unauthorized"), nil)
			ctx.Abort()
			return
		} else {
			//set cache ,next
			mcache.MixinId = mixinId
			mcache.MixinAuth = true
			err1 = util.Rdb.Replace(uid, *mcache, -1)
			if err1 != nil {
				err = util.NewErr(errors.New("cache error"), util.ErrDataBase, "")
				util.HandleResponse(ctx, err, nil)
				return
			}
		}
		ctx.Next()
	}
}

/**
 * @Description: 生成Token,uid=github.ID
 * @param uid
 * @return string
 * @return error
 */
func GenToken(uid string) (string, error) {

	c := MyClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(util.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "sky",                                           // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)

}
/**
 * @Description: 解析jwt为Myclaim
 * @param tokenString
 * @return *MyClaims
 * @return error
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

/**
 * @Description: 判断请求的Token情况,经过该中间件验证,ctx中一定有cache的key,但是不一定授权了mixin
	使用memory做缓存时重启可能导致token合法，但是没有对应cache,自动查询数据库，填充mixin是否登录
 * @return func(c *gin.Context)
 */
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		log.Debug("authHeader = ", authHeader)

		//无Token,需要授权github
		if authHeader == "" {
			log.Debug("No Token")
			util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrUnauthorized, "request have no token"), nil)
			ctx.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.HandleResponse(ctx, util.NewErr(errors.New(""), util.ErrUnauthorized, "authorization format error"), nil)
			ctx.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err1 := ParseToken(parts[1])
		if err1 != nil {
			util.HandleResponse(ctx, util.NewErr(err1, util.ErrUnauthorized, "invalid token"), nil)
			ctx.Abort()
			return
		}

		/*
			mcache := &util.MCache{}
			err1 = util.Rdb.Get(claim.Uid,mcache)
			if err1 != nil{
				util.HandleResponse(ctx,util.NewErr(err1,util.ErrDataBase,"cache get error"),nil)
				return
			}

			if mcache.MixinAuth{
				ctx.Next()
				return
			}
			//更新mixin信息
			mixin_id,err := service.GetMixinIdByUserId(*mcache.Github.ID)
			if err != nil{
				util.HandleResponse(ctx,err,nil)
				ctx.Abort()
				return
			}

			if mixin_id == ""{
				ctx.Next()
				return
			}else {
				//set cache ,next
				mcache.MixinId = mixin_id
				mcache.MixinAuth = true
				err1 = util.Rdb.Replace(claim.Uid,*mcache,-1)
				if err1 != nil{
					err = util.NewErr(errors.New("cache error"), util.ErrDataBase, "")
					util.HandleResponse(ctx, err, nil)
					return
				}
			}
		*/

		//set Key
		ctx.Set(util.UID, claim.Uid)
		//uid = randomUid不是githubId
		ctx.Next()
	}
}
