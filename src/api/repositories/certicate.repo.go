package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CertificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) *CertificateRepository {
	return &CertificateRepository{db}
}

func (repo *CertificateRepository) Update(id uuid.UUID, model *models.CertificateEntity) (*models.CertificateEntity, error) {
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

func (repo *CertificateRepository) Insert(model *models.CertificateEntity) (*models.CertificateEntity, error) {
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

func (repo *CertificateRepository) GetAll() ([]models.CertificateEntity, error) {
	var err error
	var model []models.CertificateEntity
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
