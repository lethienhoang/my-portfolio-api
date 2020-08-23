package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type EducationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) *EducationRepository {
	return &EducationRepository{db}
}

func (repo *EducationRepository) Update(id uuid.UUID, model *models.EducationEntity) (*models.EducationEntity, error) {
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

func (repo *EducationRepository) Insert(model *models.EducationEntity) (*models.EducationEntity, error) {
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

func (repo *EducationRepository) GetAll() ([]models.EducationEntity, error) {
	var err error
	var model []models.EducationEntity
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
