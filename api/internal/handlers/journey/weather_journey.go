package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	weatherParser "github.com/Central-University-IT-prod/backend-oleg-top/api/external/weather"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func WeatherJourney(c fiber.Ctx) error {
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
		Order("start_date ASC").
		Model(&models.Point{}).
		Where("journey_name = ?", journey.Name).
		Find(&points)

	type response struct {
		Point   models.Point            `json:"point"`
		Weather []weatherParser.Weather `json:"weather"`
	}

	var res []response

	wp := weatherParser.NewWeatherParser()

	for _, p := range points {
		wp.SetPeriod(p.StartDate, p.EndDate)
		wp.SetLocation(p.Location)

		res = append(res, response{
			Point:   p,
			Weather: wp.ShowWeather(),
		})
	}

	return c.JSON(res)
}
