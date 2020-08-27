package routes

import (
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
func Load(db *gorm.DB) []Route {
	skillRoutes := GetSkillRoutes(db)
	profileRotues := GetProfileRoutes(db)
	educationRotues := GetEducationRoutes(db)
	certificateRotues := GetCertificateRoutes(db)

	routes := skillRoutes
	routes = append(routes, profileRotues...)
	routes = append(routes, educationRotues...)
	routes = append(routes, certificateRotues...)

	return routes
}

// SetupRoutesWithMiddleware setup middleware for routes or otherwise
func SetupRoutesWithMiddleware(db *gorm.DB, g *gin.Engine) {
	for _, route := range Load(db) {
		if route.AuthRequired {
			authorized := g.Group("/")
			// define authen here
			authorized.Use(..) {
				authorized.Handle(route.method, route.URI, route.Handler)
			}
		} 
		else {
			g.Handle(route.method, route.URI, route.Handler)
		}
	}
}
