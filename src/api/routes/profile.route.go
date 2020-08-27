package routes

import (
	"my-portfolio-api/controllers"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetProfileRoutes(db *gorm.DB) []Route {
	controller := controllers.NewProfileController(db)

	routes := []Route{
		Route{
			URI:          "/profile?type={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/profile?manufacturer={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/profile",
			Method:       http.Get,
			Handler:      controller.GetSkills,
			AuthRequired: false,
		},
	}

	return routes
}
