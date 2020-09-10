package services

import (
	"my-portfolio-api/models"
	"my-portfolio-api/utils/redisserver"
	"time"

	uuid "github.com/satori/go.uuid"
)

// StoreTokenToRedis store access token or refresh token to redis
func StoreTokenToRedis(userid uuid.UUID, token *models.Token) error {
	at := time.Unix(token.AtExpires, 0)
	now := time.Now()

	_, err := redisserver.REDISCLIENT.Set(token.AccessUUID, userid.String(), at.Sub(now)).Result()
	if err != nil {
		return err
	}

	return nil
}

// DeleteTokenFromRedis remove access token when user logout
func DeleteTokenFromRedis(accessTokenUUID uuid.UUID) error {
	_, err := redisserver.REDISCLIENT.Del(accessTokenUUID.String()).Result()
	if err != nil {
		return err
	}

	return nil
}
