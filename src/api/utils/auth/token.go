package auth

import (
	"fmt"
	"my-portfolio-api/config"
	"my-portfolio-api/models"
	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/redisserver"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GenerateJWT creates a new token to the client
func GenerateJWT(userid uuid.UUID) (*models.Token, error) {
	td := &models.Token{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires

	var err error
	at := jwt.NewWithClaims(jwt.SigningMethodPS256, atClaims)
	td.AccessToken, err = at.SignedString(config.SECRETKEY)
	if err != nil {
		return td, err
	}

	return td, nil
}

// ExtractToken retrieves the token from headers ans Query
func ExtractToken(r *gin.Context) (*jwt.Token, error) {
	token, err := request.ParseFromRequestWithClaims(r.Request, request.AuthorizationHeaderExtractor, &models.Claim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.SECRETKEY, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenValid check token validation
func TokenValid(r *gin.Context) (models.Claim, error) {
	token, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(models.Claim)
	if !ok && !token.Valid {
		return nil, err
	}

	_, err := redisserver.REDISCLIENT.Get(claim.AccessTokenUUID).Result()
	if err != nil {
		return nil, err
	}

	return claim, nil
}

// ClaimFromContext reviced claim from context
func ClaimFromContext(r *gin.Context) models.Claim {
	value, _ := r.Request.Context().Value(auth.UserKey("user")).(models.Claim)
	return value
}
