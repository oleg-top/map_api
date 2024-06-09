package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	hotelParser "github.com/Central-University-IT-prod/backend-oleg-top/api/external/hotel"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

// TODO: journeyName -> points -> points_hotels -> return
func HotelsJourney(c fiber.Ctx) error {
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
		Point  models.Point        `json:"point"`
		Hotels []hotelParser.Hotel `json:"hotels"`
	}

	var res []response

	hp := hotelParser.NewHotelParser()

	for _, p := range points {
		hp.SetLocation(p.Location)
		hp.SetCheckIn(p.StartDate)
		hp.SetCheckOut(p.EndDate)
		hp.SetLimit(7)

		temp := response{
			Point:  p,
			Hotels: hp.ShowHotels(),
		}

		res = append(res, temp)
	}

	return c.JSON(res)
}
