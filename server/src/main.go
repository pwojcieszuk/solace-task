package main

import (
	"os"
	"server/database"

	"github.com/gofiber/fiber/v2"
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
	handler.SetupRoutes(app)
	handler.InitDb()
	return app
}

func main() {
	app := CreateApp(&AppHandlerStruct{})
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}
