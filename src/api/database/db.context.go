package database

import (
	"fmt"
	"my-portfolio-api/models"

	"github.com/jinzhu/gorm"
)

// DbContext is the struct
type DbContext struct {
	db *gorm.DB
}

// NewDbContext is initialize, passing db params and return db
func NewDbContext(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*DbContext, error) {
	DBURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &DbContext{
		db: db,
	}, nil
}

// GetDbContext is context of db conenction
func (d *DbContext) GetDbContext() *gorm.DB {
	return d.db
}

// Close closes connection to db
func (d *DbContext) Close() error {
	return d.db.Close()
}

// Automigrate migariton model schema
func (d *DbContext) Automigrate() error {
	return d.db.AutoMigrate(&models.SkillEntity{}, &models.ProfileEntity{}, &models.EducationEntity{}, &models.CertificateEntity{}).Error
}
