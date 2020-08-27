package controllers

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"

	"github.com/jinzhu/gorm"
)

type EducationController struct {
	educationService *services.EducationService
}

func NewEducationController(db *gorm.DB) *EducationController {
	repo := repositories.NewEducationRepository(db)
	service := services.NewEducationService(repo)

	return &EducationController{
		educationService: service,
	}
}
