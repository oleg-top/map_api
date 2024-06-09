package travelerHandler

import (
	"encoding/json"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func AddTraveler(c fiber.Ctx) error {
	var req models.Traveler

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	var journey models.Journey
	result := database.DB.Where("name = ?", req.JourneyName).First(&journey)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с указанным именем не найдено",
		})
	}

	var chat models.Chat
	result = database.DB.Where("id = ?", req.ChatID).First(&chat)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Чат с указанным идентификатором не найден",
		})
	}

	database.DB.Create(&req)

	var travelers []models.Traveler
	database.DB.Where("journey_name = ?", req.JourneyName).Find(&travelers)

	return c.Status(201).JSON(travelers)
}
