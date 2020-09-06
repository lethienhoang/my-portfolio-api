package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll from the DB
func GetEducations(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()

	repo := repositories.NewEducationRepository(dbContext.GetDbContext())
	service := services.NewEducationService(repo)

	results, err := service.GetEducations()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err.Error()
	}

	responses.OK(c, results)
}
