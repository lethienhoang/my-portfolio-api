package services

import (
	"errors"
	"my-portfolio-api/repositories"
	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/security"

	"github.com/gin-gonic/gin"
)

// UserService contain UserRepository
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService is initialize
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

// SignIn is login
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

// SignOut is logout system
func (sv *UserService) SignOut(c *gin.Context) error {
	var err error
	claim := auth.ClaimFromContext(c)

	if err = DeleteTokenFromRedis(claim.AccessTokenUUID); err != nil {
		return err
	}

	return nil
}
