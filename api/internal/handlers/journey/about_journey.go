package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func AboutJourney(c fiber.Ctx) error {
	journeyName := c.Params("journey_name")

	var journey models.Journey
	result := database.DB.Where("name = ?", journeyName).First(&journey)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с указанным названием не найдено",
		})
	}

	type response struct {
		Journey   models.Journey    `json:"journey"`
		Route     []models.Point    `json:"route"`
		Travelers []models.Traveler `json:"travelers"`
	}

	var points []models.Point
	database.DB.Model(&models.Point{}).Order("start_date ASC").Where("journey_name = ?", journeyName).Find(&points)

	var travelers []models.Traveler
	database.DB.Model(&models.Traveler{}).Where("journey_name = ?", journeyName).Find(&travelers)

	res := response{
		Journey:   journey,
		Route:     points,
		Travelers: travelers,
	}

	return c.JSON(res)
}
