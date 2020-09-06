package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

func GetCertificateRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/certificates",
			Method:       http.Get,
			Handler:      controllers.GetCertificates,
			AuthRequired: false,
		},
	}

	return routes
}
