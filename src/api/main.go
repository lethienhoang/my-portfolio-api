package main

import (
	"fmt"
	"log"
	"my-portfolio-api/config"
	"my-portfolio-api/database"
	_ "my-portfolio-api/docs"
	"my-portfolio-api/middlewares"
	"my-portfolio-api/routes"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.Load()
	// redisserver.Load()
}

// Run message
func Run() {
	g := gin.New()
	v1 := g.Group("/v1/api")
	Middleware(v1)
	Db()

	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	Listen(config.PORT, g)
}

// Listen run route, port
func Listen(port int, g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run(fmt.Sprintf(":%d", port))
}

// Db run migration when app is started
func Db() {
	db, _ := database.ConnectDb()

	var err error
	err = db.Migration()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}

//Middleware use middleware in pipline of each request
func Middleware(g *gin.RouterGroup) {
	var err error
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	var log = logrus.New()
	logFileName := time.Now().UTC().Format("2006-01-02") + "-log.log"
	fileName := path.Join(dir, "/logs/"+logFileName)

	var f *os.File

	f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error: ", err)
	}

	log.SetOutput(f)
	g.Use(middlewares.LoggerMiddleware(log))
	g.Use(middlewares.CORSMiddleware())
	g.Use(middlewares.JSONMiddlware())
	routes.SetupRoutesWithMiddleware(g)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /v1/api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	Run()
}
