package services

import (
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"

	uuid "github.com/satori/go.uuid"
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

// UpdateProfile update data to db
func (sv *ProfileService) UpdateProfile(id uuid.UUID, model *models.ProfileEntity) (*models.ProfileEntity, error) {
	entity, err := sv.profileRepo.Update(id, model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// InsertProfile insert data to db
func (sv *ProfileService) InsertProfile(model *models.ProfileEntity) (*models.ProfileEntity, error) {
	entity, err := sv.profileRepo.Insert(model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
