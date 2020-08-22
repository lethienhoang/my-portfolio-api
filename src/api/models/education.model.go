package models

import "time"

type EducationEntity struct {
	BaseEntity
	Name         string    `gorm:"size:30"`
	StartDate    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EndDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	FieldOfStudy string    `gorm:"size:30"`
	Degree       string    `gorm:"size:50"`
	Grade        string    `gorm:"size:15"`
}
