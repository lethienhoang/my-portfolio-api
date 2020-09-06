package middlewares

import (
	"context"
	"log"
	"net/http"

	"my-portfolio-api/utils/auth"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware displays a info message of the API
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s%s %s", c.Request.Method, c.Request.Host, c.Request.RequestURI, c.Request.Proto)
		c.Next()
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
			auth.UserKey("user"),
			claim)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
