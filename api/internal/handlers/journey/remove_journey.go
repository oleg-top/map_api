package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func RemoveJourney(c fiber.Ctx) error {
	journeyName := c.Params("journey_name")

	database.DB.Where("journey_name = ?", journeyName).Delete(&models.Traveler{}, &models.JourneyNote{}, &models.Point{})
	database.DB.Where("name = ?", journeyName).Delete(&models.Journey{})

	return c.JSON(fiber.Map{
		"message": "Путешествие было успешно удалено",
	})
}
