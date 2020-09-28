package database

import (
	"database/sql"
	"log"
	"my-portfolio-api/config"
	"my-portfolio-api/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbContext is the struct
type DbContext struct {
	db *gorm.DB
}

// ConnectDb connect database
func ConnectDb() (*DbContext, error) {
	sqlStr, err := sql.Open(config.DBDRIVER, config.DBURL)

	db, err := gorm.Open(postgres.New(postgres.Config{
		// DSN:        config.DBURL,
		// DriverName: config.DBDRIVER,
		Conn: sqlStr,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

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

// Automigrate migariton model schema
func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.SkillEntity{},
		&models.ProfileEntity{},
		&models.EducationEntity{},
		&models.CertificateEntity{},
		&models.UserEntity{},
		&models.SkillEntity{})
}
