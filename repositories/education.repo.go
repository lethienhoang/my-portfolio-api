package repositories

import (
	"my-portfolio-api/infrastructures"
	"my-portfolio-api/models"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type EducationRepository struct {
	db *gorm.DB
}

func NewEducationRepository() *EducationRepository {
	var dbContext infrastructures.DbContext
	return &EducationRepository{dbContext.GetDbContext()}
}

func (repo *EducationRepository) Update(id uuid.UUID, model *models.EducationEntity) (*models.EducationEntity, error) {
	err := repo.db.Model(&model).Where("Id=?", id).Update(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *EducationRepository) Insert(model *models.EducationEntity) (*models.EducationEntity, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *EducationRepository) GetAll() ([]models.EducationEntity, error) {
	var model []models.EducationEntity

	err := repo.db.Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
