package middlewares

import (
	"context"
	"time"

	"my-portfolio-api/utils/auth"
	"my-portfolio-api/utils/https"
	"my-portfolio-api/utils/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// LoggerMiddleware logs a gin HTTP request in JSON format
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	logger.SetLevel(log.ErrorLevel)
	logger.SetFormatter(&log.TextFormatter{})

	return func(c *gin.Context) {
		// Start timer
		start := time.Now().UTC()

		// log.Printf("%s %s %s %s", c.Request.Method, c.Request.Host, c.Request.RequestURI, c.Request.Proto)
		// Process Request
		c.Next()

		if status := c.Writer.Status(); status != 200 {
			entry := logger.WithFields(log.Fields{
				"client_ip":  https.GetClientIP(c),
				"duration":   start,
				"method":     c.Request.Method,
				"path":       c.Request.RequestURI,
				"status":     c.Writer.Status(),
				"referrer":   c.Request.Referer(),
				"request_id": c.Writer.Header().Get("X-Request-Id"),
				"user_id":    https.GetUserID(c),
			})

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

// CORSMiddleware enables data transfer from different domains
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// AuthMiddlware authorize an access
func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := auth.TokenValid(c)
		if ok != nil {
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User needs to be signed in to access this service"})
			return
		}

		ctx := context.WithValue(
			c.Request.Context(),
			types.UserKey("user"),
			claim)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
