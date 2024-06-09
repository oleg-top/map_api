package router

import (
	helloRoutes "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/routes/hello"
	journeyRoutes "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/routes/journey"
	profileRoutes "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/routes/profile"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/uploads", "./uploads")
	app.Static("/staticmaps", "./staticmaps")

	api := app.Group("/api", logger.New())

	helloRoutes.SetupHelloRoutes(api)
	profileRoutes.SetupRegisterRoutes(api)
	journeyRoutes.SetupJourneyRoutes(api)
}
