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

func (base *BaseEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

func (base *BaseEntity) BeforeSave() {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
}

func (base *BaseEntity) BeforeUpdate() {
	base.UpdatedAt = time.Now()
}
