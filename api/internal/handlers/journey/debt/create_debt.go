package debtHandler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type CreateDebtRequest struct {
	UserChatID  int      `json:"user_chat_id"`
	Owes        []string `json:"owes"`
	Amount      int      `json:"amount"`
	Date        string   `json:"date"`
	JourneyName string   `json:"journey_name"`
}

func CreateDebt(c fiber.Ctx) error {
	var req CreateDebtRequest

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	date, err := time.Parse("02.01.2006", req.Date)
	if err != nil {
		return err
	}

	for _, owe := range req.Owes {
		result := database.DB.Model(&models.Chat{}).Where("username = ?", owe).First(&models.Chat{})

		if result.Error != nil {
			return c.Status(404).JSON(fiber.Map{
				"message": fmt.Sprintf("Пользователь %s не найден. Запрос невалиден, операция отменилась", owe),
			})
		}
	}

	for _, owe := range req.Owes {
		var chat models.Chat

		database.DB.Where("username = ?", owe).First(&chat)

		database.DB.Create(&models.Debt{
			UserChatID:  req.UserChatID,
			OweChatID:   chat.ID,
			Amount:      req.Amount / len(req.Owes),
			Date:        date,
			JourneyName: req.JourneyName,
			IsPaid:      false,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Долги успешно созданы",
	})
}
