package controllers

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"

	"github.com/jinzhu/gorm"
)

type ProfileController struct {
	profileService *services.ProfileService
}

func NewProfileController(db *gorm.DB) *ProfileController {
	repo := repositories.NewProfileRepository(db)
	service := services.NewProfileService(repo)

	return &ProfileController{
		profileService: service,
	}
}
