package controllers

import (
	"errors"
	"my-portfolio-api/database"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetSkills godoc
// @Description get all skills
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills [get]
func GetSkills(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	skills, err := service.GetSkills()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, skills)
}

// GetSkillByType godoc
// @Description get all skills by type
// @Accept  json
// @Produce  json
// @Param type path string true "Type"
// @Success 200 {array} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills/{type} [get]
func GetSkillByType(c *gin.Context) {
	param, ok := c.GetQuery("type")
	if ok == false {
		responses.ERROR(c, http.StatusBadRequest, errors.New("Skill type is requried"))
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.GetSkillByType(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// GetSkillsByManufacturer godoc
// @Description get all skills by manufacturer
// @Accept  json
// @Produce  json
// @Param manufacturer path string true "Manufacturer"
// @Success 200 {array} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills/{manufacturer} [get]
func GetSkillsByManufacturer(c *gin.Context) {
	param, ok := c.GetQuery("manufacturer")
	if ok == false {
		responses.ERROR(c, http.StatusBadRequest, errors.New("manufacturer type is requried"))
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.GetSkillsByManufacturer(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// Update godoc
// @Description get all skills by manufacturer
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param req body models.SkillEntity true "Skill"
// @Success 200 {array} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills/{:id} [put]
func Update(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	var req models.SkillEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.Update(id, &req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// Insert godoc
// @Description insertl skill to database
// @Accept  json
// @Produce  json
// @Param req body models.SkillEntity true "Skill"
// @Success 200 {object} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills [post]
func Insert(c *gin.Context) {
	var req models.SkillEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.Insert(&req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// BulkInsert godoc
// @Description insert skills to database
// @Accept  json
// @Produce  json
// @Param req body []models.SkillEntity true "Skill"
// @Success 200 {array} models.SkillEntity
// @Failure 422 {object} map[string]string
// @Router /skills/bulk [put]
func BulkInsert(c *gin.Context) {
	var req []models.SkillEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewSkillRepository(dbContext.GetDbContext())
	service := services.NewSkillService(repo)

	results, err := service.BulkInsert(req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}
