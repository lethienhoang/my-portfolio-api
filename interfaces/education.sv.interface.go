package interfaces

import (
	"my-portfolio-api/viewmodels"
)

type IEducationService interface {
	GetEducation() ([]viewmodels.EducationVM, error)
}
