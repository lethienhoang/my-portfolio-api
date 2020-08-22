package repositories

import (
	"my-portfolio-api/infrastructures"
	"my-portfolio-api/models"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository() *ProfileRepository {
	var dbContext infrastructures.DbContext
	return &ProfileRepository{dbContext.GetDbContext()}
}

func (repo *ProfileRepository) Update(id uuid.UUID, model *models.ProfileEntity) (*models.ProfileEntity, error) {
	err := repo.db.Model(&model).Where("Id=?", id).Update(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *ProfileRepository) Insert(model *models.ProfileEntity) (*models.ProfileEntity, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *ProfileRepository) Get() (*models.ProfileEntity, error) {
	var model models.ProfileEntity

	err := repo.db.Take(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}
