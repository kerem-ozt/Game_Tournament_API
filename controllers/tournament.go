package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kerem-ozt/GoodBlast_API/models"
	"github.com/kerem-ozt/GoodBlast_API/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewTournament godoc
// @Summary      Create Tournament
// @Description  creates a new tournament
// @Tags         tournaments
// @Accept       json
// @Produce      json
// @Param        req  body      models.TournamentRequest true "Tournament Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /tournaments [post]
// @Security     ApiKeyAuth
func CreateNewTournament(c *gin.Context) {
	var requestBody models.TournamentRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	participantIDs := make([]string, len(requestBody.Participants))
	for i, participantID := range requestBody.Participants {
		participantIDs[i] = participantID.Hex()
	}

	var participantObjectIDs []primitive.ObjectID
	for _, id := range participantIDs {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			response.Message = err.Error()
			response.SendResponse(c)
			return
		}
		participantObjectIDs = append(participantObjectIDs, objectID)
	}

	tournament, err := services.CreateTournament(participantObjectIDs)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"tournament": tournament}
	response.SendResponse(c)
}

// GetTournaments godoc
// @Summary      Get Tournaments
// @Description  gets tournaments with pagination
// @Tags         tournaments
// @Accept       json
// @Produce      json
// @Param        page  query    string  false  "Switch page by 'page'"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /tournaments [get]
// @Security     ApiKeyAuth
func GetTournaments(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	pageQuery := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageQuery)
	limitQuery := c.DefaultQuery("limit", "5")
	limit, _ := strconv.Atoi(limitQuery)

	tournaments, _ := services.GetTournaments(page, limit)

	hasPrev := page > 0
	hasNext := len(tournaments) > limit

	if hasNext {
		tournaments = tournaments[:limit]
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"tournaments": tournaments, "prev": hasPrev, "next": hasNext}
	response.SendResponse(c)
}

func GetTournamentById(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	tournamentIDStr := c.Query("tournamentID")
	tournamentID, err := primitive.ObjectIDFromHex(tournamentIDStr)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	tournament, err := services.GetTournamentById(tournamentID)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"tournament": tournament}
	response.SendResponse(c)
}

func ProgressTournament(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	tournamentIDStr := c.Query("tournamentID")
	tournamentID, err := primitive.ObjectIDFromHex(tournamentIDStr)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	winners, err := services.ProgressTournament(tournamentID)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	// Convert winners slice to a slice of strings
	winnersStr := make([]string, len(winners))
	for i, winner := range winners {
		winnersStr[i] = winner.Hex()
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"winners": winnersStr}
	response.SendResponse(c)
}
