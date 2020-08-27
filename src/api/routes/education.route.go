package routes

import (
	"my-portfolio-api/controllers"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetEducationRoutes(db *gorm.DB) []Route {
	controller := controllers.NewEducationController(db)

	routes := []Route{
		Route{
			URI:          "/educations?type={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/educations?manufacturer={type}",
			Method:       http.Get,
			Handler:      controller.GetSkillsByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/educations",
			Method:       http.Get,
			Handler:      controller.GetSkills,
			AuthRequired: false,
		},
	}

	return routes
}
