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

// GetProfile godoc
// @Description get my profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProfileEntity
// @Failure 422 {object} map[string]string
// @Router /profile [get]
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
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, results)
}
