# my-portfolio-api
build portfolio api for myself

# controllers
controllers folder hosts all the structs under controllers namespace, controllers are the handler of all requests coming in, to the router, its doing just that, business logic and data access layer should be done separately.

# database
Database folder, consists of setup for the system to connect to external data source, it is used to host things like database connection configurations, MySQL, MariaDB, MongoDB, DynamoDB.

# models
models is a struct reflecting our data object from / to database. models should only define data structs, no other functionalities should be included here.

# repositories
repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.

# services
services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.

# viewmodels
viewmodels folder hosts all the structs under viewmodels namespace, viewmodels represent model struct to be use as a response return of REST API call

# routes
route.go is where we binds controllers to appropriate route to handle desired http request.

# untils
channel.go is where we handle result of goroutine for repository.
reponse.go is where we handler response of API

# main.go
main.go is the entry point of our system, initial db connection, routes, middlware for entire application here lies the router bindings it triggers Gin router singleton and call InitRouter to bind the router.
