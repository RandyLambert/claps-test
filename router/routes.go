package router

import (
	"claps-test/controller"
	"claps-test/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//添加日志中间件
	r.Use(middleware.LoggerToFile())
	//设置创建基于cookie的存储引擎,secret是加密的秘钥
	store := cookie.NewStore([]byte("secret11111"))
	//注册session中间件,设置session的sssion的名字,也是cookie的key
	r.Use(sessions.Sessions("SessionId", store))

	// /api
	apiGroup := r.Group("/api")
	{
		// /api/authinfo
		apiGroup.GET("/authInfo", controller.AuthInfo)

		// /api/oauth
		apiGroup.GET("/oauth", controller.Oauth)

		//给项目捐赠
		//https://claps.dev/api/bots/469e9ddc-25b3-35f0-8e43-17ffa80963c2/assets/c6d0c728-2624-429b-8e0d-d9d19b6592fa
		apiGroup.GET("bots/:botId/assets/:assetId", controller.Bot)

		// /api/projects
		projectsGroup := apiGroup.Group("projects")
		{
			projectsGroup.GET("/", controller.Projects)
			projectsGroup.GET("/:name", controller.Project)
			projectsGroup.GET("/:name/members", controller.ProjectMembers)
			projectsGroup.GET("/:name/transactions", controller.ProjectTransactions)

		}

		// /api/mixin
		mixinGroup := apiGroup.Group("/mixin")
		{
			mixinGroup.GET("/assets", controller.MixinAssets)
			mixinGroup.GET("/oauth", middleware.GithubAuthMiddleware(), controller.MixinOauth)
		}

		// /api/user
		userGroup := apiGroup.Group("/user")
		userGroup.Use(middleware.GithubAuthMiddleware())
		{
			userGroup.GET("/profile", controller.UserProfile)
			userGroup.GET("/assets", controller.UserAssets)
			//userGroup.GET("/transactions", controller.UserTransactions)
			//查询所有完成和未完成的记录
			userGroup.GET("/transfers", controller.UserTransfer)
			//请求获得某个用户的捐赠信息的汇总,包括总金额和捐赠人数
			userGroup.GET("/donation", controller.UserDonation)
			//提现
			userGroup.GET("/withdraw", middleware.MixinAuthMiddleware(), controller.UserWithdraw)
		}

	}

	return r
}
