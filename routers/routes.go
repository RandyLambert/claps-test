package routers

import (
	"claps-test/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) * gin.Engine{
	r.GET("/",controllers.Hello)
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/oauth",controllers.Oauth)

		projectsGroup := apiGroup.Group("projects")
		{
			projectsGroup.GET("/",controllers.Progects)
			projectsGroup.GET("/:name",controllers.Project)
			projectsGroup.GET("/:name/members",controllers.ProgectMembers)
			projectsGroup.GET("/:name/transactions",controllers.ProgectTransactions)
			//projectsGroup.GET("/:name/donations",controllers.ProgectMembers)
		}


		apiGroup.GET("/authInfo",controllers.AuthInfo)

		mixinGroup := apiGroup.Group("/mixin")
		{
			mixinGroup.GET("/assets",controllers.Assets)
		}


		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/profile")
			userGroup.GET("/assets")
			userGroup.GET("/transactions")
		}
	}

	return r
}
