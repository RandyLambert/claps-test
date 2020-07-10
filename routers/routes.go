package routers

import (
	"claps-test/common"
	"claps-test/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) * gin.Engine{
	//添加日志中间件
	r.Use(common.LoggerToFile())
	//测试
	r.GET("/",controllers.Hello)

	// /api
	apiGroup := r.Group("/api")
	{
		// /api/oauth
		apiGroup.GET("/oauth",controllers.Oauth)

		// /api/projects
		projectsGroup := apiGroup.Group("projects")
		{
			projectsGroup.GET("/",controllers.Progects)
			projectsGroup.GET("/:name",controllers.Project)
			projectsGroup.GET("/:name/members",controllers.ProgectMembers)
			projectsGroup.GET("/:name/transactions",controllers.ProgectTransactions)
			//projectsGroup.GET("/:name/donations",controllers.ProgectMembers)
		}

		// /api/authinfo
		apiGroup.GET("/authInfo",common.AuthInfo)

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
