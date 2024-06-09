package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	restaurantParser "github.com/Central-University-IT-prod/backend-oleg-top/api/external/restaurant"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func RestaurantJourney(c fiber.Ctx) error {
	journeyName := c.Params("journey_name")

	var journey models.Journey

	result := database.DB.
		Model(&journey).
		Where("name = ?", journeyName).
		First(&journey)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с указанным названием не найдено",
		})
	}

	var points []models.Point

	database.DB.
		Model(&models.Point{}).
		Where("journey_name = ?", journey.Name).
		Find(&points)

	type response struct {
		Point       models.Point                  `json:"point"`
		Restaurants []restaurantParser.Restaurant `json:"restaurants"`
	}

	var res []response

	rp := restaurantParser.NewRestaurantParser()

	for _, p := range points {
		rp.SetLocation(p.Location)

		res = append(res, response{
			Point:       p,
			Restaurants: rp.ShowRestaurants(),
		})
	}

	return c.JSON(res)
}
