package main

import (
	"os"
	"server/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type AppHandler interface {
	SetupRoutes(app *fiber.App)
	InitDb()
}

type AppHandlerStruct struct{}

func (a *AppHandlerStruct) SetupRoutes(app *fiber.App) {
	setupRoutes(app)
}

func (a *AppHandlerStruct) InitDb() {
	database.ConnectDb()
}

func CreateApp(handler AppHandler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	handler.SetupRoutes(app)
	handler.InitDb()
	return app
}

func main() {
	app := CreateApp(&AppHandlerStruct{})
	port := os.Getenv("SERVER_PORT")

	app.Listen(":" + port)
}
