package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kerem-ozt/GoodBlast_API/models"
	db "github.com/kerem-ozt/GoodBlast_API/models/db"
	"github.com/kerem-ozt/GoodBlast_API/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Progress godoc
// @Summary      Progress
// @Description  progress users score
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        req  body      models.TournamentRequest true "Progress Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /user [post]
// @Security     ApiKeyAuth
func Progress(c *gin.Context) {
	var requestBody models.ProgressRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	userID := requestBody.UserID

	err := services.Progress(userID, requestBody.Score)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.SendResponse(c)
}

func WhoAmI(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	token := c.GetHeader("Bearer-Token")
	tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
	if err != nil {
		models.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := services.FindUserById(tokenModel.User)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"user": user}
	response.SendResponse(c)
}

func Attend(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	token := c.GetHeader("Bearer-Token")
	tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
	if err != nil {
		models.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := services.FindUserById(tokenModel.User)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	tournamentIDStr := c.Query("tournamentID")
	tournamentID, err := primitive.ObjectIDFromHex(tournamentIDStr)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	err = services.Attend(user.ID, tournamentID)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.SendResponse(c)
}
