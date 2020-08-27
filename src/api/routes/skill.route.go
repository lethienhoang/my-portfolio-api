package routes

import (
	"my-portfolio-api/controllers"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetSkillRoutes(db *gorm.DB) []Route {
	controller := controllers.NewSkillController(db)

	routes := []Route{
		Route{
			URI:          "/skills?type={type}",
			Method:       http.Get,
			Handler:      controller.GetAllByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills?manufacturer={type}",
			Method:       http.Get,
			Handler:      controller.GetAllByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills",
			Method:       http.Get,
			Handler:      controller.GetAll,
			AuthRequired: false,
		},
	}

	return routes
}
