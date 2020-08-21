package interfaces

import "my-portfolio-api/viewmodels"

type ISkillService interface {
	GetSkillByType(string) ([]viewmodels.SkillVM, error)
	GetSkillsByManufacturer(string) ([]viewmodels.SkillVM, error)
	GetSkills() ([]viewmodels.SkillVM, error)
}
