package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
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
