package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

// EducationService respresents certificate repo, service
type EducationService struct {
	educationRepo repositories.EducationRepository
}

// GetEducations is getting all certificate in db
func (sv *EducationService) GetEducations() ([]viewmodels.EducationVM, error) {
	eduVM := []viewmodels.EducationVM{}

	entity, err := sv.educationRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("education is not found")
	}

	copier.Copy(&eduVM, &entity)

	return eduVM, nil
}
