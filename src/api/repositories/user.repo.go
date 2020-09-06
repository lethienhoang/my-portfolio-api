package repositories

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/channels"

	"github.com/jinzhu/gorm"
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

		err = repo.db.Where("email=?", email).Take(&model).Error
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
