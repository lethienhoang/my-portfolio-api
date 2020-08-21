package repositories

import (
	"my-portfolio-api/infrastructures"
	"my-portfolio-api/interfaces"
	"my-portfolio-api/models"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CertificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) *CertificateRepository {
	var dbContext infrastructures.DbContext
	return &CertificateRepository{dbContext.GetDbContext()}
}

var _ interfaces.ICertificateRepository = &CertificateRepository{}

func (repo *CertificateRepository) Update(id uuid.UUID, model *models.CertificateEntity) (*models.CertificateEntity, error) {
	err := repo.db.Model(&model).Where("Id=?", id).Update(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *CertificateRepository) Insert(model *models.CertificateEntity) (*models.CertificateEntity, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *CertificateRepository) GetCertificates() (*models.CertificateEntity, error) {
	var model models.CertificateEntity

	err := repo.db.Find(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}
