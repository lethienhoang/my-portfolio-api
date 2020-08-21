package viewmodels

import (
	uuid "github.com/satori/go.uuid"
)

type SkillVM struct {
	ID           uuid.UUID
	SkillName    string
	Level        int
	Experiences  float32
	Manufacturer string
}
