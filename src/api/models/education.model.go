package models

import (
	"errors"
	"time"
)

// EducationEntity validates the inputs
type EducationEntity struct {
	BaseEntity
	Name         string    `gorm:"size:30"`
	StartDate    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EndDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	FieldOfStudy string    `gorm:"size:30"`
	Degree       string    `gorm:"size:50"`
	Grade        string    `gorm:"size:15"`
}

// Validate validates the inputs
func (e *EducationEntity) Validate() error {
	if e.Name == "" {
		return errors.New("Name is required")
	}

	if e.StartDate.IsZero() {
		return errors.New("StartDate is required")
	}

	return nil
}
