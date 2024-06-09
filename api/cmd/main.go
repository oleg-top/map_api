package main

import (
	"log/slog"
	"os"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		slog.Error("missed SERVER_ADDRESS env (export smth like '0.0.0.0:8080')")
		os.Exit(1)
	}

	app := fiber.New()

	database.ConnectToDB()

	router.SetupRoutes(app)

	app.Listen(":8080")
}
