package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"

	uuid "github.com/satori/go.uuid"
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

// UpdateEducations update data into db
func (sv *EducationService) UpdateEducations(id uuid.UUID, model *models.EducationEntity) (*models.EducationEntity, error) {
	entity, err := sv.educationRepo.Update(id, model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// InsertEducations insert data into db
func (sv *EducationService) InsertEducations(model *models.EducationEntity) (*models.EducationEntity, error) {
	entity, err := sv.educationRepo.Insert(model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
