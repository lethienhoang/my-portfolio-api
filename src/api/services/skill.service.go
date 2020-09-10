package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
)

// SkillService respresents skill repo, service
type SkillService struct {
	skillRepo *repositories.SkillRepository
}

// NewSkillService is initialize
func NewSkillService(newRepo *repositories.SkillRepository) *SkillService {
	return &SkillService{
		skillRepo: newRepo,
	}
}

// GetSkillByType is getting data base on type
// type :
// * - cloud
// * - back end
// * - font end
func (sv *SkillService) GetSkillByType(skillType string) ([]models.SkillEntity, error) {
	if skillType == "" {
		return nil, errors.New("skill type is requried")
	}

	entity, err := sv.skillRepo.GetByType(skillType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	return entity, nil
}

// GetSkillsByManufacturer is getting data base on manufacturer
// * - microsoft
// * - Google
func (sv *SkillService) GetSkillsByManufacturer(manufacturerType string) ([]models.SkillEntity, error) {
	entity, err := sv.skillRepo.GetByManufacturer(manufacturerType)
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	return entity, nil
}

// GetSkills is getting all certificate in db
func (sv *SkillService) GetSkills() ([]models.SkillEntity, error) {
	entity, err := sv.skillRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entity) == 0 {
		return nil, errors.New("skill is not found")
	}

	return entity, nil
}
