package models

import "errors"

// ProfileEntity model
type ProfileEntity struct {
	BaseEntity
	FullName    string `json:"full_name"; gorm:"size:20;not null;"`
	Level       string `json:"level"; gorm:"size:20;not null;"`
	ImgURL      string `json:"img_url"; gorm:"size:255"`
	Description string `json:"description"; gorm:"size:255"`
	Gender      string `json:"gender"; gorm:"size:10"`
	Email       string `json:"email"; gorm:"size:20"`
}

// Validate validates the inputs
func (p *ProfileEntity) Validate() error {
	if p.Email == "" {
		return errors.New("Email is required")
	}

	if p.Gender == "" {
		return errors.New("Gender is required")
	}

	if p.FullName == "" {
		return errors.New("FullName is required")
	}

	if p.Description == "" {
		return errors.New("Description is required")
	}

	return nil
}
