package journeyHandler

import (
	"encoding/json"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type CreateJourneyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ChatID      int    `json:"chat_id"`
}

func CreateJourney(c fiber.Ctx) error {
	var req CreateJourneyRequest

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	var journey models.Journey
	result := database.DB.Model(&journey).Where("name = ?", req.Name).First(&journey)

	if result.Error == nil {
		return c.Status(403).JSON(fiber.Map{
			"message": "Путешествие с таким названием уже существует",
		})
	}

	var chat models.Chat
	result = database.DB.Model(&chat).Where("id = ?", req.ChatID).First(&chat)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Чат с указанным идентификатором не найден",
		})
	}

	journey.Description = req.Description
	journey.Name = req.Name

	database.DB.Create(&journey)

	database.DB.Create(&models.Traveler{
		ChatID:      req.ChatID,
		JourneyName: req.Name,
	})

	return c.Status(201).JSON(fiber.Map{
		"message": "Путешествие было успешно создано",
	})
}
