package noteHandler

import (
	"encoding/json"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func CreateNote(c fiber.Ctx) error {
	var note models.JourneyNote

	if err := json.Unmarshal(c.Body(), &note); err != nil {
		return err
	}

	var journey models.Journey
	result := database.DB.Where("name = ?", note.JourneyName).First(&journey)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с указанным названием не найдено",
		})
	}

	database.DB.Create(&note)

	return c.Status(201).JSON(fiber.Map{
		"message": "Заметка была успешно создана",
		"note":    note,
	})
}
