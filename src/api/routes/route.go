package routes

import (
	"my-portfolio-api/middlewares"

	"github.com/gin-gonic/gin"
)

// Route contain route fields
type Route struct {
	URI          string
	Method       string
	Handler      gin.HandlerFunc
	AuthRequired bool
}

// Load collect all routes
func Load() []Route {
	skillRoutes := SkillRoutes()
	profileRoutes := ProfileRoutes()
	educationRoutes := EducationRoutes()
	certificateRoutes := CertificateRoutes()
	userRoutes := UserRoutes()
	healthCheckRoutes := HealthCheckRoutes()

	routes := healthCheckRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, skillRoutes...)
	routes = append(routes, profileRoutes...)
	routes = append(routes, educationRoutes...)
	routes = append(routes, certificateRoutes...)

	return routes
}

// SetupRoutesWithMiddleware setup middleware for routes or otherwise
func SetupRoutesWithMiddleware(g *gin.RouterGroup) {
	for _, route := range Load() {
		g.Use(middlewares.LoggerMiddleware())
		g.Use(middlewares.JSONMiddlware())

		if route.AuthRequired {
			authorized := g.Group("/")
			// define authen here
			authorized.Use(middlewares.AuthMiddlware())
			{
				authorized.Handle(route.Method, route.URI, route.Handler)
			}
		} else {
			g.Handle(route.Method, route.URI, route.Handler)
		}
	}
}
