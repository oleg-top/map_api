package journeyHandler

import (
	"fmt"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/external/staticmap"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func StaticMapJourney(c fiber.Ctx) error {
	journeyName := c.Params("journey_name")
	chatID := c.Params("chat_id")

	result := database.DB.Where("name = ?", journeyName).First(&models.Journey{})

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с указанным названием не найдено",
		})
	}

	var chat models.Chat
	result = database.DB.Where("id = ?", chatID).First(&chat)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Чат с указанным идентификатором не найден",
		})
	}

	var points []models.Point
	database.DB.
		Model(&models.Point{}).
		Order("start_date ASC").
		Where("journey_name = ?", journeyName).
		Find(&points)

	if len(points) == 0 {
		return c.JSON(fiber.Map{
			"message": "В вашем путешествии еще нет локаций",
		})
	}

	path := fmt.Sprintf("./staticmaps/%s-%s.png", journeyName, chatID)

	var locations []string

	locations = append(locations, chat.Location)

	for _, p := range points {
		locations = append(locations, p.Location)
	}

	staticMapRenderer := staticmap.NewStaticMapRenderer()
	staticMapRenderer.SetSize(1920, 1080)
	staticMapRenderer.SetPath(path)
	staticMapRenderer.SetLocations(locations)

	staticMapRenderer.SaveImage()

	return c.SendFile(path)
}
