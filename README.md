# my-portfolio-api
build portfolio api for myself


# controllers
controllers folder hosts all the structs under controllers namespace, controllers are the handler of all requests coming in, to the router, its doing just that, business logic and data access layer should be done separately.

controller struct implement services through their interface, no direct services implementation should be done in controller, this is done to maintain decoupled systems. The implementation will be injected during the compiled time.

# infrasctructures
infrasctructures folder host all structs under infrasctructures namespace, infrasctructures consists of setup for the system to connect to external data source, it is used to host things like database connection configurations, MySQL, MariaDB, MongoDB, DynamoDB.

# interfaces
interfaces folder hosts all the structs under interfaces namespace, interfaces as the name suggest are the bridge between different domain so they can interact with each other, in our case, this should be the only way for them to interact.

interface in Go is a bit different then you might already find in other language like Java or C#, while the later implements interface explicitly, Go implements interface implicitly. You just need to implement all method the interface has, and you’re good to “Go”.

In our system, our PlayerController implements IPlayerService to be able to interact with the implementation that will be injected. In our case, IPlayerService will be injected with PlayerService.

The same thing applies on PlayerService which implements IPlayerRepository to be able interact with the injected implementation. In our case, IPlayerRepository will be injected with PlayerRepository during the compile time.

PlayerRepository on the other hand, will be injected with infrasctructure configuration that has been setup earlier, this ensure that you can change the implementation of PlayerRepository, without changing the implementor which in this case PlayerService let alone break it. The same thing goes to PlayerService and PlayerController relationship, we can refactor PlayerService, we can change it however we want, without touching the implementor which is PlayerController.

# models
models folder hosts all structs under models namespace, model is a struct reflecting our data object from / to database. models should only define data structs, no other functionalities should be included here.

# repositories
repositories folder hosts all structs under repositories namespace, repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.

# services
services folder hosts all structs under services namespace, services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.

controller might implement many services interface to satisfy the request needs, and controller should be agnostic of how services implements their logic, all they care is that they should be able to pull the result they need according to the interface they implements.

# viewmodels
viewmodels folder hosts all the structs under viewmodels namespace, viewmodels are model to be use as a response return of REST API call

# main.go
main.go is the entry point of our system, here lies the router bindings it triggers ChiRouter singleton and call InitRouter to bind the router.

# router.go
router.go is where we binds controllers to appropriate route to handle desired http request. By default we are using Chi router as it is a light weight router and not bloated with unnecessary unwanted features.

# servicecontainer.go
servicecontainer.go is where the magic begins, this is the place where we injected all implementations of interfaces. Lets cover throughly in the dependency injection section.