package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

// SkillService respresents skill repo, service
type SkillService struct {
	skillRepo repositories.SkillRepository
}

// GetSkillByType is getting data base on type
// type :
// * - cloud
// * - back end
// * - font end
func (sv *SkillService) GetSkillByType(skillType string) ([]viewmodels.SkillVM, error) {
	skillVM := []viewmodels.SkillVM{}

	if skillType == "" {
		return skillVM, errors.New("skill type is requried")
	}

	entity, err := sv.skillRepo.GetByType(skillType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVM, &entity)

	return skillVM, nil
}

// GetSkillsByManufacturer is getting data base on manufacturer
// * - microsoft
// * - Google
func (sv *SkillService) GetSkillsByManufacturer(manufacturerType string) ([]viewmodels.SkillVM, error) {
	skillVM := []viewmodels.SkillVM{}

	entity, err := sv.skillRepo.GetByManufacturer(manufacturerType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVM, &entity)

	return skillVM, nil
}

// GetSkills is getting all certificate in db
func (sv *SkillService) GetSkills() ([]viewmodels.SkillVM, error) {
	entity, err := sv.skillRepo.GetAll()
	if err != nil {
		return nil, err
	}

	skillVM := []viewmodels.SkillVM{}
	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVM, &entity)

	return skillVM, nil
}
