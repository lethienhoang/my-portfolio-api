package middlewares

import (
	"fmt"
	"my-portfolio-api/utils/https"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Logger logs a gin HTTP request in JSON format
func Logger() gin.HandlerFunc {
	logFileName := time.Now().UTC().Format("yyyy-MM-dd") + "-log.log"
	fileName := path.Join("/", logFileName)
	// src, err := os.OpenFile(fileName, os.O_CREATE|os.O_SYNC|os.O_APPEND|os.O_RDONLY, os.ModeAppend)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err", err)
	}

	logger := log.New()
	logger.Out = src
	logger.SetLevel(log.ErrorLevel)
	logger.SetFormatter(&log.TextFormatter{})

	return func(c *gin.Context) {
		// Start timer
		start := time.Now().UTC()

		// Process Request
		c.Next()

		entry := logger.WithFields(log.Fields{
			"client_ip":  https.GetClientIP(c),
			"duration":   start,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    https.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() != 200 {
			entry.Error(c.Errors.String())
		}
	}
}
