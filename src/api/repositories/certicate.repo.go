package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// CertificateRepository is the struct for Certificate bussiness
type CertificateRepository struct {
	db *gorm.DB
}

// NewCertificateRepository is initialize, passing database connection and return a new repo with connection
func NewCertificateRepository(db *gorm.DB) *CertificateRepository {
	return &CertificateRepository{db}
}

// Update updates a Certificate from the DB and returns a new Certificate updated or an error
func (repo *CertificateRepository) Update(id uuid.UUID, model *models.CertificateEntity) (*models.CertificateEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.CertificateEntity{}).Where("id=?", id).First(&model).Error
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

// Insert returns a new Certificate created or an error
func (repo *CertificateRepository) Insert(model *models.CertificateEntity) (*models.CertificateEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.CertificateEntity{}).Create(&model).Error
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

// GetAll returns all the Certificate from the DB
func (repo *CertificateRepository) GetAll() ([]models.CertificateEntity, error) {
	var err error
	var model []models.CertificateEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.CertificateEntity{}).Find(&model).Error
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
