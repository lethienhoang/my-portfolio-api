package models

type SkillEntity struct {
	BaseEntity
	Name        string `gorm:"size:255"`
	Level       int    `gorm:"size:100"`
	Experiences float32
	Type        string `gorm:"size:20"`
}
