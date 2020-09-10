package controllers

import (
	"my-portfolio-api/utils/responses"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	res := map[string]string{
		"status": "UP",
	}

	responses.OK(c, res)
}
