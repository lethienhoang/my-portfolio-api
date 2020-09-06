package models

import (
	"errors"
	"my-portfolio-api/utils/security"
	"strings"

	"github.com/badoux/checkmail"
)

// User model
type UserEntity struct {
	BaseEntity
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:60;not null" json:"password,omitempty"`
}

// BeforeSave hash the user password
func (u *UserEntity) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// Validate validates the inputs
func (u *UserEntity) Validate(action string) error {
	{
		switch strings.ToLower(action) {
		case "login":
			if u.Email == "" {
				return errors.New("Email is required")
			}

			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid emai")
			}

			if u.Password == "" {
				return errors.New("Password is required")
			}
		default:
			return nil
		}
	}

	return nil
}
