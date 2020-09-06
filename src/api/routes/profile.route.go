package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

func GetProfileRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/profile",
			Method:       http.Get,
			Handler:      controllers.GetProfile,
			AuthRequired: false,
		},
	}

	return routes
}
