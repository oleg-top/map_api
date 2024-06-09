package debtHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func ListDebts(c fiber.Ctx) error {
	chatID := c.Params("chat_id")

	var chat models.Chat

	result := database.DB.Where("id = ?", chatID).First(&chat)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Чат с указанным идентификатором не найден",
		})
	}

	debts_giver := make([]models.Debt, 0)
	debts_owe := make([]models.Debt, 0)

	database.DB.Model(&models.Debt{}).Where("user_chat_id = ?", chat.ID).Find(&debts_giver)
	database.DB.Model(&models.Debt{}).Where("owe_chat_id = ?", chat.ID).Find(&debts_owe)

	return c.JSON(fiber.Map{
		"giver": debts_giver,
		"owe":   debts_owe,
	})
}
