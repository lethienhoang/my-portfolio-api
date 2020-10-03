package models

import (
	"errors"
	"time"
)

// EducationEntity validates the inputs
type EducationEntity struct {
	BaseEntity
	Name         string    `json:"name"; gorm:"size:30"`
	StartDate    time.Time `json:"start_date"; gorm:"default:CURRENT_TIMESTAMP"`
	EndDate      time.Time `json:"end_date"; gorm:"default:CURRENT_TIMESTAMP"`
	FieldOfStudy string    `json:"field_of_study"; gorm:"size:30"`
	Degree       string    `json:"degree"; gorm:"size:50"`
	Grade        string    `json:"grade"; gorm:"size:15"`
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
