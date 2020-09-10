package redisserver

import (
	"my-portfolio-api/config"

	"github.com/go-redis/redis"
)

var (
	REDISCLIENT *redis.Client
)

// Load the redis server
func Load() {
	if len(config.REDIS_DSN) == 0 {
		panic("Redis is not configuration")
	}

	REDISCLIENT = redis.NewClient(&redis.Options{
		Addr:       config.REDIS_DSN,
		Password:   config.PASSWRORD,
		MaxRetries: 2,
	})

	_, err := REDISCLIENT.Ping().Result()
	if err != nil {
		panic(err)
	}
}
