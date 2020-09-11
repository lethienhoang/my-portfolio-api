package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCertificates godoc
// @Summary Show a account
// @Description get certificates
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.CertificateEntity
// @Failure 422 {object} map[string]string
// @Router /certificates [get]
func GetCertificates(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewCertificateRepository(dbContext.GetDbContext())
	service := services.NewCertificateService(repo)

	results, err := service.GetCertificates()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, results)
}
