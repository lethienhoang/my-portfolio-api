package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

// CertificateService respresents certificate repo, service
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
func (sv *CertificateService) GetCertificates() ([]viewmodels.CertificateVM, error) {
	var cerVM []viewmodels.CertificateVM

	entity, err := sv.certificateRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("certificate is not found")
	}

	copier.Copy(&cerVM, &entity)

	return cerVM, nil
}
