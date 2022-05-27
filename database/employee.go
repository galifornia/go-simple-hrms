package database

import (
	"github.com/galifornia/go-simple-hrms/types"
	"github.com/gofiber/fiber/v2"
)

func NewEmployee(c *fiber.Ctx) error {
	return c.SendString("NewEmployee")
}

func GetEmployees(c *fiber.Ctx) error {
	var leads []types.Employee
	DB.Find(&leads).Limit(10)

	return c.JSON(leads)
}

func GetEmployee(c *fiber.Ctx) error {
	return c.SendString("GetEmployee")
}

func DeleteEmployee(c *fiber.Ctx) error {
	return c.SendString("DeleteEmployee")
}

func UpdateEmployee(c *fiber.Ctx) error {
	return c.SendString("UpdateEmployee")
}
