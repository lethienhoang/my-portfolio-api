package controllers

import (
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CertificateController struct {
	certificateService *services.CertificateService
}

func NewCertificateController(db *gorm.DB) *CertificateController {
	repo := repositories.NewCertificateRepository(db)
	service := services.NewCertificateService(repo)

	return &CertificateController{
		certificateService: service,
	}
}

func (controller *CertificateController) GetAll(c *gin.Context) {
	certificates, err := controller.certificateService.GetCertificates()
	if err != nil {
		responses.ERROR(c, http.StatusUnprocessableEntity, err)
	}

	responses.OK(c, certificates)
}
