package interfaces

import (
	"my-portfolio-api/models"

	uuid "github.com/satori/go.uuid"
)

type ISkillRepository interface {
	Update(uuid.UUID, *models.SkillEntity) (*models.SkillEntity, error)
	Insert(*models.SkillEntity) (*models.SkillEntity, error)
	GetByType(string) ([]models.SkillEntity, error)
	GetByManufacturer(string) ([]models.SkillEntity, error)
	GetAll() ([]models.SkillEntity, error)
	GetById(uuid.UUID) (*models.SkillEntity, error)
}
