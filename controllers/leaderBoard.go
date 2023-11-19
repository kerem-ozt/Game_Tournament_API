package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerem-ozt/GoodBlast_API/models"
	"github.com/kerem-ozt/GoodBlast_API/services"
)

func EnsureLeaderboardInitialized(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	leaderboard := services.EnsureLeaderboardInitialized("global")
	// leaderboard, err := services.GetLeaderboard("global")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	fmt.Println("leaderboard: ", leaderboard)
	// if len(leaderboard.Users) == 0 {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "leaderboard is empty"})
	// 	return
	// }
	response.StatusCode = http.StatusOK
	response.Success = true
	// response.Data = gin.H{"notes": notes, "prev": hasPrev, "next": hasNext}
	response.SendResponse(c)
	// c.JSON(http.StatusOK, gin.H{"message": "leaderboard is initialized"})
}

func GetGlobalLeaderboard(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// leaderboard, err := services.GetLeaderboard("global")
	leaderboard, err := services.GetGlobalLeaderboard("global")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// if len(leaderboard.Users) == 0 {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "leaderboard is empty"})
	// 	return
	// }

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"leaderboard": leaderboard}
	response.SendResponse(c)
}

func GetLeaderboardByCountry(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	country := c.Query("country")
	// leaderboardtype := c.Query("leaderboardtype")

	leaderboard, err := services.GetLeaderboardByCountry("global", country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// if len(leaderboard.Users) == 0 {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "leaderboard is empty"})
	// 	return
	// }

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"leaderboard": leaderboard}
	response.SendResponse(c)
}
