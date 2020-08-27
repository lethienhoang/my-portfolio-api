package services

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

// ProfileService respresents certificate repo, service
type ProfileService struct {
	profileRepo *repositories.ProfileRepository
}

// NewProfileService is initialize
func NewProfileService(newRepo *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo: newRepo,
	}
}

// GetProfile is getting only one
func (sv *ProfileService) GetProfile() (*viewmodels.ProfileVM, error) {
	var profileVM viewmodels.ProfileVM

	entity, err := sv.profileRepo.Get()
	if err != nil {
		return nil, err
	}

	copier.Copy(&profileVM, &entity)

	return &profileVM, nil
}
