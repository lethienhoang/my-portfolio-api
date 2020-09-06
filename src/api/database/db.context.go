package database

import (
	"log"
	"my-portfolio-api/config"
	"my-portfolio-api/models"

	"github.com/jinzhu/gorm"
)

// DbContext is the struct
type DbContext struct {
	db *gorm.DB
}

// GetDbContext connect database
func ConnectDb() (*DbContext, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.LogMode(true)

	err = Automigrate(db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DbContext{
		db: db,
	}, nil
}

// GetDbContext gets gorm.DB
func (d *DbContext) GetDbContext() *gorm.DB {
	return d.db
}

// Close closes connection to db
func (d *DbContext) Close() error {
	return d.db.Close()
}

// Automigrate migariton model schema
func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.SkillEntity{}, &models.ProfileEntity{}, &models.EducationEntity{}, &models.CertificateEntity{}).Error
}
