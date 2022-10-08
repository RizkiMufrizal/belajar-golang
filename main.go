package main

import (
	"belajar-golang/config"
	"belajar-golang/controller"
	"belajar-golang/exception"
	"belajar-golang/repository"
	"belajar-golang/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Belajar Golang
// @version 2.0
// @description Sample Golang Project use Gofiber
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	//setup configuration
	configuration := config.New()
	database := config.ConnectDb(configuration)

	// Setup Repository
	employeeRepository := repository.NewEmployeeRepository(database)

	// Setup Repository
	employeeService := service.NewEmployeeService(&employeeRepository)

	// Setup Controller
	employeeController := controller.NewEmployeeController(&employeeService)

	app := fiber.New()
	app.Use(cors.New())

	// add to route
	employeeController.Route(app)

	err := app.Listen(configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
