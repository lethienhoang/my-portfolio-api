package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetByEmail(email string) (models.UserEntity, error) {
	model := models.UserEntity{}
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.UserEntity{}).Where("email=?", email).Take(&model).Error
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

func (repo *UserRepository) Insert(model *models.UserEntity) error {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.UserEntity{}).Select("Email", "Password").Create(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return nil
	}

	return err
}

func (repo *UserRepository) ExistEmail(email string) error {
	var model models.UserEntity
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.UserEntity{}).Where("email=?", email).Take(&model).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return nil
	}

	return err
}

func (repo *UserRepository) UpdatePassword(id uuid.UUID, password string) error {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = repo.db.Model(&models.UserEntity{}).Where("id=?", id).Update("Password", password).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return nil
	}

	return err
}
