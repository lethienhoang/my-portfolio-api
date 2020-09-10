package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// UserRoutes define user routes
func UserRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/SignIn",
			Method:       http.MethodPost,
			Handler:      controllers.SignIn,
			AuthRequired: false,
		},
		Route{
			URI:          "/SignOut",
			Method:       http.MethodGet,
			Handler:      controllers.SignOut,
			AuthRequired: true,
		},
	}

	return routes
}
