package handlers

import "github.com/gofiber/fiber/v2"

func SetupAppRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")
	setupEmployeeRoutes(api)
}
