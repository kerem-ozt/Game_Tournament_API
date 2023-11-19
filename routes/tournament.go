package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/controllers"
)

func TournamentRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	tournaments := router.Group("/tournament", handlers...)
	{
		tournaments.POST(
			"/create",
			// validators.CreateTournamentValidator(),
			controllers.CreateNewTournament,
		)

		tournaments.POST(
			"/creategroup",
			// validators.CreateTournamentValidator(),
			controllers.CreateTournamentGroups,
		)

		tournaments.GET(
			"/getall",
			// validators.GetTournamentsValidator(),
			controllers.GetTournaments,
		)

		tournaments.GET(
			"/getbyid",
			// validators.GetTournamentsValidator(),
			controllers.GetTournamentById,
		)

		tournaments.POST(
			"/progress",
			// validators.CreateTournamentValidator(),
			controllers.ProgressTournament,
		)
	}
}
