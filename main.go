package main

import (
	"log"

	"github.com/galifornia/go-simple-hrms/database"
	"github.com/galifornia/go-simple-hrms/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Open & automigrate database
	if err := database.OpenDB(); err != nil {
		log.Fatal("Could not open Mongo database")
	}

	// Fiber setup
	app := fiber.New()
	handlers.SetupAppRoutes(app)

	app.Listen("localhost:3003")
}
