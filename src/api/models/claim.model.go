package models

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Claim struct {
	Authorized      bool      `json:"authorized"`
	AccessTokenUUID uuid.UUID `json:"access_uuid"`
	UserID          uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
