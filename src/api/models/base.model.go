package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        string    `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) {
	newUUID := uuid.NewV4()
	base.ID = newUUID.String()
}

func (base *BaseEntity) BeforeSave(tx *gorm.DB) {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
}

func (base *BaseEntity) BeforeUpdate(tx *gorm.DB) {
	base.UpdatedAt = time.Now()
}
