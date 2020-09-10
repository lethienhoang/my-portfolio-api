package main

import (
	"fmt"
	"my-portfolio-api/config"
	"my-portfolio-api/routes"
	"my-portfolio-api/utils/redisserver"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
	redisserver.Load()
}

// Run message
func Run() {
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	Listen(config.PORT)
}

// Listen run route, port
func Listen(port int) {
	g := gin.New()
	v1 := g.Group("/v1/api")

	routes.SetupRoutesWithMiddleware(v1)
	g.Run(fmt.Sprintf(":%d", port))
}

func main() {
	Run()
}
