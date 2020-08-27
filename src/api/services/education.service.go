package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
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
func (sv *EducationService) GetEducations() ([]viewmodels.EducationVM, error) {
	var eduVM []viewmodels.EducationVM

	entity, err := sv.educationRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return eduVM, errors.New("education is not found")
	}

	copier.Copy(&eduVM, &entity)

	return eduVM, nil
}
