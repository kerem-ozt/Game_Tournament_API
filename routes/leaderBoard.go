package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/controllers"
)

func leaderboardRouter(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	leaderBoard := router.Group("/leaderBoard", handlers...)
	{
		leaderBoard.GET(
			"",
			// validators.CreateTournamentValidator(),
			controllers.EnsureLeaderboardInitialized,
		)

		leaderBoard.GET(
			"get",
			// validators.CreateTournamentValidator(),
			controllers.GetLeaderboard,
		)
	}
}
