package routes

import (
	"my-portfolio-api/controllers"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetCertificateRoutes(db *gorm.DB) []Route {
	controller := controllers.NewCertificateController(db)

	routes := []Route{
		Route{
			URI:          "/certificates?type={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/certificates?manufacturer={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/certificates",
			Method:       http.Get,
			Handler:      controller.GetSkills,
			AuthRequired: false,
		},
	}

	return routes
}
