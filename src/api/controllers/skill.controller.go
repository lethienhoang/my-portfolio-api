package controllers

import (
	"errors"
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSkills from the DB
func GetSkills(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	skills, err := service.GetSkills()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err.Error()
	}

	responses.OK(c, skills)
}

// GetSkillByType from the DB
func GetSkillByType(c *gin.Context) {
	param, ok := c.GetQuery("type")
	if ok == false {
		responses.ERROR(c, http.StatusBadRequest, errors.New("Skill type is requried"))
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.GetSkillByType(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err.Error()
	}

	responses.OK(c, results)
}

// GetSkillsByManufacturer from the DB
func GetSkillsByManufacturer(c *gin.Context) {
	param, ok := c.GetQuery("manufacturer")
	if ok == false {
		responses.ERROR(c, http.StatusBadRequest, errors.New("manufacturer type is requried"))
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.GetSkillsByManufacturer(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err.Error()
	}

	responses.OK(c, results)
}
