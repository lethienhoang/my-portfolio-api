package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// UserRoutes define user routes
func UserRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "users/signin",
			Method:       http.MethodPost,
			Handler:      controllers.SignIn,
			AuthRequired: false,
		},
		Route{
			URI:          "users/signout",
			Method:       http.MethodGet,
			Handler:      controllers.SignOut,
			AuthRequired: true,
		},
		Route{
			URI:          "users/update-password",
			Method:       http.MethodPut,
			Handler:      controllers.UpdatePassword,
			AuthRequired: true,
		},
		Route{
			URI:          "users/signup",
			Method:       http.MethodPost,
			Handler:      controllers.SignUp,
			AuthRequired: false,
		},
	}

	return routes
}
