package routes

import (
	"my-portfolio-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Route struct {
	URI          string
	method       string
	Handler      *gin.HandlerFunc
	AuthRequired bool
}

// Load collect all routes
func Load() []Route {
	skillRoutes := GetSkillRoutes()
	profileRotues := GetProfileRoutes()
	educationRotues := GetEducationRoutes()
	certificateRotues := GetCertificateRoutes()

	routes := skillRoutes
	routes = append(routes, profileRotues...)
	routes = append(routes, educationRotues...)
	routes = append(routes, certificateRotues...)

	return routes
}

// SetupRoutesWithMiddleware setup middleware for routes or otherwise
func SetupRoutesWithMiddleware(g *gin.Engine) {
	for _, route := range Load() {
		g.Use(middlewares.LoggerMiddleware)
		g.Use(middlewares.JSONMiddlware)

		if route.AuthRequired {
			authorized := g.Group("/")
			// define authen here
			authorized.Use(middlewares.AuthMiddlware) {
				authorized.Handle(route.method, route.URI, route.Handler)
			}
		} 
		else {
			g.Handle(route.method, route.URI, route.Handler)
		}
	}
}
