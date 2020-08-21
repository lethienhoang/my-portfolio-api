package interfaces

import (
	"my-portfolio-api/models"

	uuid "github.com/satori/go.uuid"
)

type IProfileRepository interface {
	Update(uuid.UUID, *models.ProfileEntity) (*models.ProfileEntity, error)
	Insert(*models.ProfileEntity) (*models.ProfileEntity, error)
	Get() (*models.ProfileEntity, error)
}
