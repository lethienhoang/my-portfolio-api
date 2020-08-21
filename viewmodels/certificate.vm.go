package viewmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CertificateVM struct {
	ID        uuid.UUID
	Name      string
	IssueDate time.Time
	URL       string
	Authority string
}
