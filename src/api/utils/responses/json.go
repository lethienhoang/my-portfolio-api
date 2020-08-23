package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK represents a response successfully
func OK(r *gin.Context, statusCode int, data interface{}) {
	r.JSON(http.StatusOK, data)
}

// ERROR represents a response failure
func ERROR(r *gin.Context, statusCode int, err error) {
	r.JSON(statusCode, gin.H{
		"errors": string(err.Error()),
	})
}
