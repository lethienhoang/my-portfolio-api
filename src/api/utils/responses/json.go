package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(r *gin.Context, statusCode int, data interface{}) {
	r.JSON(http.StatusOK, data)
}

func ERROR(r *gin.Context, statusCode int, err error) {
	r.JSON(statusCode, gin.H{
		"errors": string(err.Error()),
	})
}
