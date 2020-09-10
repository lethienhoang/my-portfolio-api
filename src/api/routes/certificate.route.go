package routes

import (
	"my-portfolio-api/controllers"
	"net/http"
)

// CertificateRoutes define certificate routes
func CertificateRoutes() []Route {
	routes := []Route{
		Route{
			URI:          "/certificates",
			Method:       http.MethodGet,
			Handler:      controllers.GetCertificates,
			AuthRequired: false,
		},
	}

	return routes
}
