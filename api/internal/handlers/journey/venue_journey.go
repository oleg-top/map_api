package journeyHandler

import (
	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	venueParser "github.com/Central-University-IT-prod/backend-oleg-top/api/external/venue"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func VenueJourney(c fiber.Ctx) error {
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
		Venues []venueParser.Venue `json:"venues"`
	}

	var res []response

	vp := venueParser.NewVenueParser()

	for _, p := range points {
		vp.SetLocation(p.Location)

		res = append(res, response{
			Point:  p,
			Venues: vp.ShowVenues(),
		})
	}

	return c.JSON(res)
}
