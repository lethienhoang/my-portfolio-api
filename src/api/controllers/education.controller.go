package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
		return
	}

	repo := repositories.NewEducationRepository(dbContext.GetDbContext())
	service := services.NewEducationService(repo)

	results, err := service.GetEducations()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// UpdateEducations godoc
// @Description get educations
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param req body models.EducationEntity true "Education"
// @Success 200 {object} models.EducationEntity
// @Failure 422 {object} map[string]string
// @Router /educations/{:id} [put]
func UpdateEducations(c *gin.Context) {
	var req models.EducationEntity
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	id, _ := uuid.FromString(c.Param("id"))

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewEducationRepository(dbContext.GetDbContext())
	service := services.NewEducationService(repo)

	results, err := service.UpdateEducations(id, &req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// InsertEducations godoc
// @Description get educations
// @Accept  json
// @Produce  json
// @Param req body models.EducationEntity true "Education"
// @Success 200 {object} models.EducationEntity
// @Failure 422 {object} map[string]string
// @Router /educations [post]
func InsertEducations(c *gin.Context) {
	var req models.EducationEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewEducationRepository(dbContext.GetDbContext())
	service := services.NewEducationService(repo)

	results, err := service.InsertEducations(&req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}
