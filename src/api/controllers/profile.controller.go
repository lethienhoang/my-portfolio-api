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

// UpdateProfile godoc
// @Description get my profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProfileEntity
// @Failure 422 {object} map[string]string
// @Router /profile/{:id} [put]
func UpdateProfile(c *gin.Context) {
	var err error
	var id uuid.UUID

	var req models.ProfileEntity
	if err = c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
	}

	id, err = uuid.FromString(c.Query("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewProfileRepository(dbContext.GetDbContext())
	service := services.NewProfileService(repo)

	results, err := service.UpdateProfile(id, &req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, results)
}

// InsertProfile godoc
// @Description get my profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProfileEntity
// @Failure 422 {object} map[string]string
// @Router /profile [post]
func InsertProfile(c *gin.Context) {
	var req models.ProfileEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewProfileRepository(dbContext.GetDbContext())
	service := services.NewProfileService(repo)

	results, err := service.InsertProfile(&req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, results)
}
