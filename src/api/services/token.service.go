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

	status := redisserver.REDISCLIENT.Set(token.AccessUUID, userid, at.Sub(now))
	if status.Err != nil {
		return status.Err()
	}

	return nil
}

func DeleteTokenFromRedis(userid string) error {
	deleted := redisserver.REDISCLIENT.Del(userid)
	if deleted.Err != nil {
		return deleted.Err()
	}

	return nil
}
