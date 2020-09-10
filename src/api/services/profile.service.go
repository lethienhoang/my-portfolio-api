package services

import (
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
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
func (sv *ProfileService) GetProfile() (*models.ProfileEntity, error) {
	entity, err := sv.profileRepo.Get()
	if err != nil {
		return nil, err
	}

	return entity, nil
}
