package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/controllers"
)

func UserRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/user", handlers...)
	{
		user.GET(
			"/whoami",
			// validators.CreateUserValidator(),
			controllers.WhoAmI,
		)

		user.GET(
			"/getall",
			// validators.CreateUserValidator(),
			controllers.GetAllUsers,
		)

		user.GET(
			"/getbyid",
			// validators.CreateUserValidator(),
			controllers.GetById,
		)

		user.DELETE(
			"/delete",
			// validators.CreateUserValidator(),
			controllers.DeleteUser,
		)

		user.POST(
			"/attendtotournament",
			// validators.GetUsersValidator(),
			controllers.AttendToTournament,
		)

		user.POST(
			"/updatestat",
			// validators.GetUserValidator(),
			controllers.UpdateUserStat,
		)
	}
}
