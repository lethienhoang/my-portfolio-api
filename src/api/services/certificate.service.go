package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"

	uuid "github.com/satori/go.uuid"
)

// CertificateService contain CertificateRepository
type CertificateService struct {
	certificateRepo *repositories.CertificateRepository
}

// NewCertificateService is initialize
func NewCertificateService(newRepo *repositories.CertificateRepository) *CertificateService {
	return &CertificateService{
		certificateRepo: newRepo,
	}
}

// GetCertificates is getting all certificate in db
func (sv *CertificateService) GetCertificates() ([]models.CertificateEntity, error) {
	entity, err := sv.certificateRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("certificate is not found")
	}

	return entity, nil
}

// UpdateCertificates update data in db
func (sv *CertificateService) UpdateCertificates(id uuid.UUID, model *models.CertificateEntity) (*models.CertificateEntity, error) {
	entity, err := sv.certificateRepo.Update(id, model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// InsertCertificates insert data into db
func (sv *CertificateService) InsertCertificates(model *models.CertificateEntity) (*models.CertificateEntity, error) {
	entity, err := sv.certificateRepo.Insert(model)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
