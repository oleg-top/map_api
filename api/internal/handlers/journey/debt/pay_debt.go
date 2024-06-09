package debtHandler

import (
	"encoding/json"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type PayDebtRequest struct {
	UserChatID  int    `json:"user_chat_id"`
	Owe         string `json:"owe"`
	Amount      int    `json:"amount"`
	JourneyName string `json:"journey_name"`
}

func PayDebt(c fiber.Ctx) error {
	var req PayDebtRequest

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	var owe models.Chat
	result := database.DB.Where("username = ?", req.Owe).First(&owe)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Пользователь не найден",
		})
	}

	var debt models.Debt
	result = database.DB.
		Where("user_chat_id = ? AND owe_chat_id = ? AND amount = ? AND journey_name = ?",
			req.UserChatID, owe.ID, req.Amount, req.JourneyName).
		First(&debt)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Долг не найден",
		})
	}

	debt.IsPaid = true

	database.DB.Save(&debt)

	return c.JSON(fiber.Map{
		"message": "Долг был успешно выплачен",
	})
}
