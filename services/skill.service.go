package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

type SkillService struct {
	skillRepo repositories.SkillRepository
}

func (sv *SkillService) GetSkillByType(skillType string) ([]viewmodels.SkillVM, error) {
	skillVm := []viewmodels.SkillVM{}

	if skillType == "" {
		return skillVm, errors.New("skill type is requried")
	}

	entity, err := sv.skillRepo.GetByType(skillType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVm, &entity)

	return skillVm, nil
}

func (sv *SkillService) GetSkillsByManufacturer(manufacturerType string) ([]viewmodels.SkillVM, error) {
	skillVm := []viewmodels.SkillVM{}

	entity, err := sv.skillRepo.GetByManufacturer(manufacturerType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVm, &entity)

	return skillVm, nil
}

func (sv *SkillService) GetSkills() ([]viewmodels.SkillVM, error) {
	entity, err := sv.skillRepo.GetAll()
	if err != nil {
		return nil, err
	}

	skillVm := []viewmodels.SkillVM{}
	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	copier.Copy(&skillVm, &entity)

	return skillVm, nil
}
