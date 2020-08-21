package viewmodels

import (
	uuid "github.com/satori/go.uuid"
)

type ProfileVM struct {
	ID          uuid.UUID
	FullName    string
	Level       string
	ImgURL      string
	Description string
	Gender      string
	Email       string
}
