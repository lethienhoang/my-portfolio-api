package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP";swaggerignore:"true"`
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) {
	uuid := uuid.NewV4()
	base.ID = uuid
}

func (base *BaseEntity) BeforeSave(tx *gorm.DB) {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
}

func (base *BaseEntity) BeforeUpdate(tx *gorm.DB) {
	base.UpdatedAt = time.Now()
}
