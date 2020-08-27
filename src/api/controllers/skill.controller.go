package controllers

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SkillController struct {
	skillService *services.SkillService
}

func NewSkillController(db *gorm.DB) *SkillController {
	skillRepo := repositories.NewSkillRepository(db)
	skillService := services.NewSkillService(skillRepo)

	return &SkillController{
		skillService: skillService,
	}
}

func (controller *SkillController) GetAll(c *gin.Context) {
	skills, err := controller.skillService.GetSkills()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, skills)
}

func (controller *SkillController) GetAllByType(c *gin.Context) {
	param, isErr := c.GetQuery("type")
	if isErr {
		responses.ERROR(c, http.StatusBadRequest, errors.New("Skill type is requried"))
	}

	skills, err := controller.skillService.GetSkillByType(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, skills)
}

func (controller *SkillController) GetAllByManufacturer(c *gin.Context) {
	param, isErr := c.GetQuery("manufacturer")
	if isErr {
		responses.ERROR(c, http.StatusBadRequest, errors.New("manufacturer type is requried"))
	}

	skills, err := controller.skillService.GetSkillsByManufacturer(param)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, skills)
}
