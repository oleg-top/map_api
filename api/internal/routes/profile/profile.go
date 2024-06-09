package profileRoutes

import (
	profileHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/profile"
	"github.com/gofiber/fiber/v3"
)

func SetupRegisterRoutes(router fiber.Router) {
	profileRouter := router.Group("/profile")
	profileRouter.Post("/update", profileHandler.UpdateProfile)
	profileRouter.Get("/about/:chat_id", profileHandler.AboutProfile)
}
