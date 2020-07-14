package routers

import (
	"claps-test/common"
	"claps-test/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) * gin.Engine{
	//添加日志中间件
	r.Use(common.LoggerToFile())
	//测试
	r.GET("/",controllers.Hello)
	//设置创建基于cookie的存储引擎,secret是加密的秘钥
	store := cookie.NewStore([]byte("secret11111"))
	//注册session中间件,设置session的sssion的名字,也是cookie的key
	r.Use(sessions.Sessions("SessionId", store))

	// /api
	apiGroup := r.Group("/api")
	{
		// /api/authinfo
		apiGroup.GET("/authInfo",common.AuthInfo)

		// /api/oauth
		apiGroup.GET("/oauth",controllers.Oauth)

		// /api/projects
		projectsGroup := apiGroup.Group("projects")
		{
			projectsGroup.GET("/",controllers.Projects)
			projectsGroup.GET("/:name",controllers.Project)
			projectsGroup.GET("/:name/members",controllers.ProjectMembers)
			projectsGroup.GET("/:name/transactions",controllers.ProjectTransactions)
			//projectsGroup.GET("/:name/donations",controllers.ProgectMembers)
		}


		// /api/mixin
		mixinGroup := apiGroup.Group("/mixin")
		{
			mixinGroup.GET("/assets",controllers.Assets)
		}


		// /api/user
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/profile",controllers.Profile)
			userGroup.GET("/assets")
			userGroup.GET("/transactions")
		}
	}

	return r
}