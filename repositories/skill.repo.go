package repositories

import (
	"my-portfolio-api/infrastructures"
	"my-portfolio-api/interfaces"
	"my-portfolio-api/models"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SkillRepository struct {
	db *gorm.DB
}

func NewSkillRepository() *SkillRepository {
	var dbContext infrastructures.DbContext
	return &SkillRepository{dbContext.GetDbContext()}
}

var _ interfaces.ISkillRepository = &SkillRepository{}

func (repo *SkillRepository) Update(id uuid.UUID, model *models.SkillEntity) (*models.SkillEntity, error) {
	err := repo.db.Model(&model).Where("Id=?", id).Update(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *SkillRepository) Insert(model *models.SkillEntity) (*models.SkillEntity, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *SkillRepository) GetByType(skillType string) ([]models.SkillEntity, error) {
	var model []models.SkillEntity

	err := repo.db.Where("Type=?", skillType).Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *SkillRepository) GetByManufacturer(manufacturerType string) ([]models.SkillEntity, error) {
	var model []models.SkillEntity

	err := repo.db.Where("Manufacturer=?", manufacturerType).Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *SkillRepository) GetAll() ([]models.SkillEntity, error) {
	var model []models.SkillEntity

	err := repo.db.Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *SkillRepository) GetById(id uuid.UUID) (*models.SkillEntity, error) {
	var model models.SkillEntity

	err := repo.db.Take(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}
