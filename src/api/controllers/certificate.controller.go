package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/models"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
		return
	}

	defer dbContext.Close()

	repo := repositories.NewCertificateRepository(dbContext.GetDbContext())
	service := services.NewCertificateService(repo)

	results, err := service.GetCertificates()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}

// UpdateCertificates godoc
// @Summary Show a account
// @Description get certificates
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param req body models.CertificateEntity true "Certificate"
// @Success 200 {object} models.CertificateEntity
// @Failure 422 {object} map[string]string
// @Router /certificates/{:id} [put]
func UpdateCertificates(c *gin.Context) {
	var req models.CertificateEntity
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	id, _ := uuid.FromString(c.Query("id"))

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewCertificateRepository(dbContext.GetDbContext())
	service := services.NewCertificateService(repo)

	results, err := service.UpdateCertificates(id, &req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	defer dbContext.Close()

	responses.OK(c, results)
}

// InsertCertificates godoc
// @Summary Show a account
// @Description get certificates
// @Accept  json
// @Produce  json
// @Param req body models.CertificateEntity true "Certificate"
// @Success 200 {object} models.CertificateEntity
// @Failure 422 {object} map[string]string
// @Router /certificates [post]
func InsertCertificates(c *gin.Context) {
	var req models.CertificateEntity
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	defer dbContext.Close()

	repo := repositories.NewCertificateRepository(dbContext.GetDbContext())
	service := services.NewCertificateService(repo)

	results, err := service.InsertCertificates(&req)
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}

	responses.OK(c, results)
}
