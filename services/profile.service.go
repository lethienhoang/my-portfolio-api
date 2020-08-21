package services

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/viewmodels"

	"github.com/jinzhu/copier"
)

type ProfileService struct {
	profileRepo repositories.ProfileRepository
}

func (sv *ProfileService) GetProfile() (*viewmodels.ProfileVM, error) {
	profileVm := viewmodels.ProfileVM{}

	entity, err := sv.profileRepo.Get()
	if err != nil {
		return nil, err
	}

	copier.Copy(&profileVm, &entity)

	return &profileVm, nil
}
