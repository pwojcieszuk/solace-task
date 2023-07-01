package main

import (
	"github.com/gofiber/fiber/v2"

	"server/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/health", handlers.Health)

}
