package services

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

// ProfileService respresents certificate repo, service
type ProfileService struct {
	profileRepo repositories.ProfileRepository
}

// GetProfile is getting only one
func (sv *ProfileService) GetProfile() (*viewmodels.ProfileVM, error) {
	profileVM := viewmodels.ProfileVM{}

	entity, err := sv.profileRepo.Get()
	if err != nil {
		return nil, err
	}

	copier.Copy(&profileVM, &entity)

	return &profileVM, nil
}
