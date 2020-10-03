package models

import "errors"

// SkillEntity model
type SkillEntity struct {
	BaseEntity
	Name         string  `json:"name"; gorm:"size:255"`
	Level        int     `json:"level"; gorm:"size:100"`
	Experiences  float32 `json:"experiences"`
	Type         string  `json:"type"; gorm:"size:20"`
	Manufacturer string  `json:"manfacturer"; gorm:"size:25"`
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
