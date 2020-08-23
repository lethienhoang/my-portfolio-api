package routes

import "github.com/gin-gonic/gin"

type Route struct {
	URI          string
	method       string
	Handler      *gin.HandlerFunc
	AuthRequired bool
}

// Load collect all routes
func Load() []Route {

}

// SetupRoutesWithMiddleware setup middleware for routes or otherwise
func SetupRoutesWithMiddleware(g *gin.Engine) {
	for _, route := range Load() {
		if route.AuthRequired {
			authorized := g.Group("/")
			// define authen here
			authorized.Use(..) {
				authorized.Handle(route.method, route.URI, route.Handler)
			}
		} else {
			g.Handle(route.method, route.URI, route.Handler)
		}
	}
}
