package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// HealthCheckRoutes define healthcheck route
func HealthCheckRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/healthcheck",
			Method:       http.MethodGet,
			Handler:      controllers.HealthCheck,
			AuthRequired: false,
		},
	}

	return routes
}
