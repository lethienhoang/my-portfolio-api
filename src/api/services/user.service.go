package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/security"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (sv *UserService) SignIn(email, password string) (string, error) {
	if email == "" {
		return "", errors.New("Email is required")
	}

	if password == "" {
		return "", errors.New("Password is required")
	}

	entity, err := sv.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if err := security.VerifyPassword(entity.Password, password); err != nil {
		return "", errors.New("Password incorrect")
	}

	token, err := auth.GenerateJWT(entity.ID)
	if err != nil {
		return "", errors.New("Generate token failed")
	}

	if err := StoreTokenToRedis(entity.ID, token); err != nil {
		return "", errors.New("Insert token failed")
	}

	return token.AccessToken, nil
}

func (sv *UserService) SignOut(userid string) error {
	var err error

	if userid == "" {
		return errors.New("Email is required")
	}

	if err = DeleteTokenFromRedis(userid); err != nil {
		return err
	}

	return nil
}
