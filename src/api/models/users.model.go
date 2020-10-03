package models

import (
	"errors"
	"my-portfolio-api/utils/security"
	"strings"
	"unicode"

	"github.com/badoux/checkmail"
)

// User model
type UserEntity struct {
	BaseEntity
	Email    string `json:"email"; gorm:"size:100;not null;unique"`
	Password string `json:"password,omitempty"; gorm:"size:60;not null"`
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

// PasswordValidate validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.
func PasswordValidate(pass string) bool {
	var (
		upp, low, num, symb bool
		length              uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			length++
		case unicode.IsLower(char):
			low = true
			length++
		case unicode.IsNumber(char):
			num = true
			length++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			symb = true
			length++
		default:
			return false
		}
	}

	if !upp || !low || !num || !symb || length < 8 {
		return false
	}

	return true
}

// Validate validates the inputs of user created
func (u *UserEntity) Validate(action string) error {
	{
		switch strings.ToLower(action) {
		case "insert":
			if u.Email == "" {
				return errors.New("Email is required")
			}

			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid emai")
			}

			if u.Password == "" {
				return errors.New("Password is required")
			}

			if ok := PasswordValidate(u.Password); ok == false {
				s := "-At least one upper case letter \n" +
					"-At least one lower case letter \n" +
					"-At least one digit \n" +
					"-At least one special character \n" +
					"-At least eight characters long"

				return errors.New(s)
			}

		case "udpate":
			if u.Password == "" {
				return errors.New("Password is required")
			}

			if ok := PasswordValidate(u.Password); ok == false {
				s := "-At least one upper case letter \n" +
					"-At least one lower case letter \n" +
					"-At least one digit \n" +
					"-At least one special character \n" +
					"-At least eight characters long"

				return errors.New(s)
			}
		default:
			return nil
		}
	}

	return nil
}
