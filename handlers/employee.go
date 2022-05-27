package handlers

import (
	"github.com/galifornia/go-simple-hrms/database"
	"github.com/gofiber/fiber/v2"
)

func setupEmployeeRoutes(app fiber.Router) {
	employee := app.Group("employee")
	employee.Get("/", database.GetEmployees)
	employee.Get("/:id", database.GetEmployee)
	employee.Post("/", database.NewEmployee)
	employee.Delete("/:id", database.DeleteEmployee)
	employee.Put("/:id", database.UpdateEmployee)
}
