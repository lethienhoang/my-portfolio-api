package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SkillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *SkillRepository {
	return &SkillRepository{db}
}

func (repo *SkillRepository) Update(id uuid.UUID, model *models.SkillEntity) (*models.SkillEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&model).Where("Id=?", id).Update(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return nil, err
}

func (repo *SkillRepository) Insert(model *models.SkillEntity) (*models.SkillEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Create(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return nil, err
}

func (repo *SkillRepository) GetByType(skillType string) ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Where("Type=?", skillType).Find(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return nil, err
}

func (repo *SkillRepository) GetByManufacturer(manufacturerType string) ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Where("Manufacturer=?", manufacturerType).Find(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return nil, err
}

func (repo *SkillRepository) GetAll() ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Find(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return nil, err
}

func (repo *SkillRepository) GetById(id uuid.UUID) (*models.SkillEntity, error) {
	var err error
	var model models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Take(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return &model, nil
	}

	return nil, err
}
