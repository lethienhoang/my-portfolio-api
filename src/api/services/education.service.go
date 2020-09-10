package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
)

// EducationService respresents certificate repo, service
type EducationService struct {
	educationRepo *repositories.EducationRepository
}

// NewEducationService is initialize
func NewEducationService(newRepo *repositories.EducationRepository) *EducationService {
	return &EducationService{
		educationRepo: newRepo,
	}
}

// GetEducations is getting all certificate in db
func (sv *EducationService) GetEducations() ([]models.EducationEntity, error) {
	entity, err := sv.educationRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("education is not found")
	}

	return entity, nil
}
