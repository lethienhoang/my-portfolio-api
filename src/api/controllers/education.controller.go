package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEducations godoc
// @Description get educations
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.EducationEntity
// @Failure 422 {object} map[string]string
// @Router /educations [get]
func GetEducations(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewEducationRepository(dbContext.GetDbContext())
	service := services.NewEducationService(repo)

	results, err := service.GetEducations()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, results)
}
