package viewmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type EducationVM struct {
	ID           uuid.UUID
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	FieldOfStudy string
	Degree       string
	Grade        string
}
