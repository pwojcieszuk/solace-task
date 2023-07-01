package main

import (
	"github.com/gofiber/fiber/v2"

	"server/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/health", handlers.Health)

	app.Get("/favorites", handlers.ListFavorites)
	app.Get("/favorites/:id", handlers.GetFavorite)
	app.Post("/favorites", handlers.CreateFavorite)
	app.Patch("/favorites/:id", handlers.UpdateFavorite)
	app.Delete("/favorites/:id", handlers.DeleteFavorite)
}
