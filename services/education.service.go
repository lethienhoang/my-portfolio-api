package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

type EducationService struct {
	educationRepo repositories.EducationRepository
}

func (sv *EducationService) GetEducations() ([]viewmodels.EducationVM, error) {
	eduVm := []viewmodels.EducationVM{}

	entity, err := sv.educationRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("education is not found")
	}

	copier.Copy(&eduVm, &entity)

	return eduVm, nil
}
