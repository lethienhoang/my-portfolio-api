package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

func GetSkillRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/skills?type={type}",
			Method:       http.Get,
			Handler:      controllers.GetSkillByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills?manufacturer={type}",
			Method:       http.Get,
			Handler:      controllers.GetSkillsByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills",
			Method:       http.Get,
			Handler:      controllers.GetSkills,
			AuthRequired: false,
		},
	}

	return routes
}
