package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileService *services.ProfileService
}

// GetProfile from the DB
func GetProfile(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewProfileRepository(dbContext.GetDbContext())
	service := services.NewProfileService(repo)

	results, err := service.GetProfile()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err.Error()
	}

	responses.OK(c, results)
}
