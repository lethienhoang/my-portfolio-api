package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK represents a response successfully
func OK(r *gin.Context, item interface{}) {
	r.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// ERROR represents a response failure
func ERROR(r *gin.Context, statusCode int, err error) {
	res := map[string]string{
		"message": string(err.Error()),
	}

	r.AbortWithStatusJSON(statusCode, gin.H{
		"data": res,
	})
}
