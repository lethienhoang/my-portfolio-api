package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// EducationRoutes define education routes
func EducationRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/educations",
			Method:       http.MethodGet,
			Handler:      controllers.GetEducations,
			AuthRequired: false,
		},
	}

	return routes
}
