package controllers

import (
	"my-portfolio-api/utils/responses"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Description check api healthcheck
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /educations [get]
func HealthCheck(c *gin.Context) {
	res := map[string]string{
		"status": "UP",
	}

	responses.OK(c, res)
}
