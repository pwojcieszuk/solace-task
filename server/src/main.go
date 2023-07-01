package main

import (
	"net/http"
	"os"
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

// TODO route for storibng favorites in format (ts) MAp<number, {title: string, image: string}>
// favorites.set(mal_id, { title: title, image: image });

func main() {
	database.ConnectDb()
	app := fiber.New()
	port := os.Getenv("PORT")
	setupRoutes(app)

	app.Listen(":" + port)
}
