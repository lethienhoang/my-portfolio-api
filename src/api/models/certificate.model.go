package models

import (
	"errors"
	"time"
)

// CertificateEntity validates the inputs
type CertificateEntity struct {
	BaseEntity
	Name      string    `gorm:"size:100"`
	IssueDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	URL       string    `gorm:"size:255"`
	Authority string    `gorm:"size:150"`
}

// Validate validates the inputs
func (c *CertificateEntity) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.IssueDate.IsZero() {
		return errors.New("IssueDate is required")
	}

	if c.URL == "" {
		return errors.New("URL is required")
	}

	if c.Authority == "" {
		return errors.New("Authority is required")
	}

	return nil
}
