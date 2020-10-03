package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// SkillRoutes define skill routes
func SkillRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/skills?type=:type",
			Method:       http.MethodGet,
			Handler:      controllers.GetSkillByType,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills?manufacturer=:type",
			Method:       http.MethodGet,
			Handler:      controllers.GetSkillsByManufacturer,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills",
			Method:       http.MethodGet,
			Handler:      controllers.GetSkills,
			AuthRequired: false,
		},
		Route{
			URI:          "/skills/:id",
			Method:       http.MethodPut,
			Handler:      controllers.Update,
			AuthRequired: true,
		},
		Route{
			URI:          "/skills",
			Method:       http.MethodPost,
			Handler:      controllers.Insert,
			AuthRequired: true,
		},
		Route{
			URI:          "/skills/bulk",
			Method:       http.MethodPost,
			Handler:      controllers.BulkInsert,
			AuthRequired: true,
		},
	}

	return routes
}
