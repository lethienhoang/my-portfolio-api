package services

import (
	"errors"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/security"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

// SignUp register user account
func (sv *UserService) SignUp(email, password string) error {
	userEntity := models.UserEntity{
		Email:    email,
		Password: password,
	}

	if err := userEntity.Validate("insert"); err != nil {
		return err
	}

	if err := sv.repo.ExistEmail(email); err != nil {
		return err
	}

	newPass, err := security.Hash(password)
	if err != nil {
		return err
	}

	userEntity.Password = string(newPass)

	if err := sv.repo.Insert(&userEntity); err != nil {
		return err
	}

	return nil
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

// UpdateUserPassword user could update own password
func (sv *UserService) UpdateUserPassword(id uuid.UUID, password string) error {
	if ok := models.PasswordValidate("update"); ok == false {
		return errors.New("Password is invalid")
	}

	newPass, err := security.Hash(password)
	if err != nil {
		return err
	}

	if err := sv.repo.UpdatePassword(id, string(newPass)); err != nil {
		return err
	}

	return nil
}
