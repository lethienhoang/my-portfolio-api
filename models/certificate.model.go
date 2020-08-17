package models

import "time"

type CertificateEntity struct {
	BaseEntity
	Name      string    `gorm:"size:100"`
	IssueDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	URL       string    `gorm:"size:255"`
	Authority string    `gorm:"size:150"`
}
