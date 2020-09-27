package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// SkillRepository is the struct for kill bussiness
type SkillRepository struct {
	db *gorm.DB
}

// NewSkillRepository is initialize, passing database connection and return a new repo with connection
func NewSkillRepository(db *gorm.DB) *SkillRepository {
	return &SkillRepository{db}
}

// Update updates a Skill from the DB and returns a new Skill updated or an error
func (repo *SkillRepository) Update(id uuid.UUID, model *models.SkillEntity) (*models.SkillEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Omit("id").Where("id=?", id).First(&model).Error
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

// Insert returns a new Skill created or an error
func (repo *SkillRepository) Insert(model *models.SkillEntity) (*models.SkillEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Create(&model).Error
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

// BatchInsert intert batch data and return array
func (repo *SkillRepository) BatchInsert(skills []models.SkillEntity) ([]models.SkillEntity, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Create(&skills).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return skills, nil
	}

	return nil, err
}

// GetByType returns Skill records base on type
func (repo *SkillRepository) GetByType(skillType string) ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Where("type=?", skillType).Find(&model).Error
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

// GetByManufacturer returns Skill records base on manufacturer
func (repo *SkillRepository) GetByManufacturer(manufacturerType string) ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Where("manufacturer=?", manufacturerType).Find(&model).Error
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

// GetAll returns all the Skill from the DB
func (repo *SkillRepository) GetAll() ([]models.SkillEntity, error) {
	var err error
	var model []models.SkillEntity
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Find(&model).Error
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

// GetByID returns only one the Profile base on id from the DB
func (repo *SkillRepository) GetByID(id uuid.UUID) (models.SkillEntity, error) {
	var err error
	model := models.SkillEntity{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.SkillEntity{}).Where("id=?", id).Take(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return model, nil
	}

	return model, err
}
