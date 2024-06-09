package profileHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func AboutProfile(c fiber.Ctx) error {
	chatID := c.Params("chat_id")

	var chat models.Chat
	result := database.DB.Model(&chat).Where("id = ?", chatID).First(&chat)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Пользователь не найден",
		})
	}

	return c.JSON(chat)
}
