package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// ProfileRoutes define profile routes
func ProfileRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/profile",
			Method:       http.MethodGet,
			Handler:      controllers.GetProfile,
			AuthRequired: false,
		},
	}

	return routes
}
