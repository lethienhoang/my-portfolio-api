package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/https"
	"my-portfolio-api/utils/types"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoggerMiddleware displays a info message of the API
// func LoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		log.Printf("%s %s%s %s", c.Request.Method, c.Request.Host, c.Request.RequestURI, c.Request.Proto)
// 		c.Next()
// 	}
// }

// LoggerMiddleware logs a gin HTTP request in JSON format
func LoggerMiddleware() gin.HandlerFunc {
	logFileName := time.Now().UTC().Format("yyyy-MM-dd") + "-log.log"
	fileName := path.Join("/logs/", logFileName)
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

		log.Printf("%s %s%s %s", c.Request.Method, c.Request.Host, c.Request.RequestURI, c.Request.Proto)
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

// JSONMiddlware set the application Content-Type
func JSONMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

// AuthMiddlware authorize an access
func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := auth.TokenValid(c)
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User needs to be signed in to access this service"})
		}

		ctx := context.WithValue(
			c.Request.Context(),
			types.UserKey("user"),
			claim)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
