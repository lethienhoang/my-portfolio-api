package interfaces

import (
	"my-portfolio-api/models"

	uuid "github.com/satori/go.uuid"
)

type ICertificateRepository interface {
	Update(uuid.UUID, *models.CertificateEntity) (*models.CertificateEntity, error)
	Insert(*models.CertificateEntity) (*models.CertificateEntity, error)
	GetCertificates() (*models.CertificateEntity, error)
}
