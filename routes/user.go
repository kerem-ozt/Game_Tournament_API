package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/controllers"
)

func UserRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/user", handlers...)
	{
		user.GET(
			"",
			// validators.CreateUserValidator(),
			controllers.WhoAmI,
		)

		user.POST(
			"/attend",
			// validators.GetUsersValidator(),
			controllers.Attend,
		)

		user.POST(
			"",
			// validators.GetUserValidator(),
			controllers.Progress,
		)
	}
}
