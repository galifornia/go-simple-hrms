package handlers

import (
	"github.com/galifornia/go-simple-hrms/database"
	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(app *fiber.App) {
	employee := app.Group("/api/v1/employee")
	employee.Get("/", database.GetEmployees)
	employee.Get("/:id", database.GetEmployee)
	employee.Post("/", database.NewEmployee)
	employee.Delete("/:id", database.DeleteEmployee)
	employee.Put("/:id", database.UpdateEmployee)
}
