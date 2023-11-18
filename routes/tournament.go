package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/controllers"
)

func TournamentRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	tournaments := router.Group("/tournaments", handlers...)
	{
		tournaments.POST(
			"",
			// validators.CreateTournamentValidator(),
			controllers.CreateNewTournament,
		)

		tournaments.GET(
			"",
			// validators.GetTournamentsValidator(),
			controllers.GetTournaments,
		)

		tournaments.POST(
			"progressTournament",
			// validators.CreateTournamentValidator(),
			controllers.ProgressTournament,
		)
	}
}
