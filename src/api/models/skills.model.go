package models

import "errors"

// SkillEntity model
type SkillEntity struct {
	BaseEntity
	Name         string `gorm:"size:255"`
	Level        int    `gorm:"size:100"`
	Experiences  float32
	Type         string `gorm:"size:20"`
	Manufacturer string `gorm:"size:25"`
}

// Validate validates the inputs
func (s *SkillEntity) Validate() error {
	if s.Name == "" {
		return errors.New("Name is required")
	}

	if s.Type == "" {
		return errors.New("Type is required")
	}

	return nil
}
