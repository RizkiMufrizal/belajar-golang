package controller

import (
	"belajar-golang/exception"
	"belajar-golang/model"
	"belajar-golang/service"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(EmployeeService *service.EmployeeService) EmployeeController {
	return EmployeeController{EmployeeService: *EmployeeService}
}

func (controller *EmployeeController) Route(app *fiber.App) {
	app.Get("/api/employee", controller.FindAll)
	app.Get("/api/employee/:id", controller.FindOne)
	app.Post("/api/employee", controller.Insert)
	app.Put("/api/employee", controller.Update)
	app.Delete("/api/employee/:id", controller.Delete)
}

// Insert func for creates a new employee.
// @Description Create a new employee.
// @Summary create a new employee
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} models.Employee
// @Router /v1/employee [post]
func (controller *EmployeeController) Insert(app *fiber.Ctx) error {
	var request model.Employee
	err := app.BodyParser(&request)
	exception.PanicIfNeeded(err)

	controller.EmployeeService.Insert(request)
	return app.JSON(model.GlobalResponse{
		Code:   200,
		Status: "OK",
		Data:   request,
	})
}

// Update func for update a employee.
// @Description Update a employee.
// @Summary Update a employee
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} models.Employee
// @Router /v1/employee [put]
func (controller *EmployeeController) Update(app *fiber.Ctx) error {
	var request model.Employee
	err := app.BodyParser(&request)
	exception.PanicIfNeeded(err)

	controller.EmployeeService.Update(request)
	return app.JSON(model.GlobalResponse{
		Code:   200,
		Status: "OK",
		Data:   request,
	})
}

// Delete func for delete a employee.
// @Description delete a employee.
// @Summary delete a employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Router /v1/employee/{id} [delete]
func (controller *EmployeeController) Delete(app *fiber.Ctx) error {
	employee := controller.EmployeeService.FindOne(app.Params("id"))
	controller.EmployeeService.Delete(employee)
	return app.JSON(model.GlobalResponse{
		Code:   200,
		Status: "OK",
		Data:   employee,
	})
}

// FindOne func for get a employee.
// @Description get a employee.
// @Summary get a employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Router /v1/employee/{id} [get]
func (controller *EmployeeController) FindOne(app *fiber.Ctx) error {
	employee := controller.EmployeeService.FindOne(app.Params("id"))
	return app.JSON(model.GlobalResponse{
		Code:   200,
		Status: "OK",
		Data:   employee,
	})
}

// FindAll func for get a employee.
// @Description get a employee.
// @Summary get a employee
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} models.Employee
// @Router /v1/employee [get]
func (controller *EmployeeController) FindAll(app *fiber.Ctx) error {
	employee := controller.EmployeeService.FindAll()
	return app.JSON(model.GlobalResponse{
		Code:   200,
		Status: "OK",
		Data:   employee,
	})
}
