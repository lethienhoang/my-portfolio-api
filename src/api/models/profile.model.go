package models

type ProfileEntity struct {
	BaseEntity
	FullName    string `gorm:"size:20;not null;"`
	Level       string `gorm:"size:20;not null;"`
	ImgURL      string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Gender      string `gorm:"size:10"`
	Email       string `gorm:"size:20"`
}
