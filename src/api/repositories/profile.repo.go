package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (repo *ProfileRepository) Update(id uuid.UUID, model *models.ProfileEntity) (*models.ProfileEntity, error) {
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

func (repo *ProfileRepository) Insert(model *models.ProfileEntity) (*models.ProfileEntity, error) {
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

func (repo *ProfileRepository) Get() (*models.ProfileEntity, error) {
	var err error
	var model models.ProfileEntity
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
