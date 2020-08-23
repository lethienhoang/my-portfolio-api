package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// EducationRepository is the struct for Education bussiness
type EducationRepository struct {
	db *gorm.DB
}

// NewEducationRepository is initialize, passing database connection and return a new repo with connection
func NewEducationRepository(db *gorm.DB) *EducationRepository {
	return &EducationRepository{db}
}

// Update updates a education from the DB and returns a new education updated or an error
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

// Insert returns a new Education created or an error
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

// GetAll returns all the Education from the DB
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
