package interfaces

import (
	"my-portfolio-api/models"

	uuid "github.com/satori/go.uuid"
)

type IEducationRepository interface {
	Update(uuid.UUID, *models.EducationEntity) (*models.EducationEntity, error)
	Insert(*models.EducationEntity) (*models.EducationEntity, error)
	GetAll() ([]models.EducationEntity, error)
}
