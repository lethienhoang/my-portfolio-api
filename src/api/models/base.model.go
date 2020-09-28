package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()
	base.ID = newUUID
	return
}

func (base *BaseEntity) BeforeSave(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return
}

func (base *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now()
	return
}
