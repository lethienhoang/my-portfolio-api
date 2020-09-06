package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

func GetEducationRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/educations",
			Method:       http.Get,
			Handler:      controllers.GetEducations,
			AuthRequired: false,
		},
	}

	return routes
}
