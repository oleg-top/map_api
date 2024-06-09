package journeyHandler

import (
	"strconv"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func ChatJourney(c fiber.Ctx) error {
	chatID := c.Params("chat_id")

	var travelers []models.Traveler
	database.DB.Model(&models.Traveler{}).Where("chat_id = ?", chatID).Find(&travelers)

	if len(travelers) == 0 {
		return c.JSON(fiber.Map{
			"message": "Вы еще не участвуете ни в одном путешествии",
		})
	}

	type response struct {
		Journey   models.Journey       `json:"journey"`
		Route     []models.Point       `json:"route"`
		Travelers []models.Traveler    `json:"travelers"`
		Notes     []models.JourneyNote `json:"notes"`
	}

	var res []response
	for _, t := range travelers {
		var j models.Journey
		database.DB.Model(&j).Where("name = ?", t.JourneyName).First(&j)

		var points []models.Point
		database.DB.Model(&models.Point{}).Order("start_date ASC").Where("journey_name = ?", t.JourneyName).Find(&points)

		var travelers []models.Traveler
		database.DB.Model(&models.Traveler{}).Where("journey_name = ?", t.JourneyName).Find(&travelers)

		var notes []models.JourneyNote
		database.DB.Model(&models.JourneyNote{}).Where("journey_name = ?", t.JourneyName).Find(&notes)

		temp := make([]models.JourneyNote, 0)

		for _, note := range notes {
			authorStr := strconv.Itoa(note.Author)

			if authorStr == chatID || note.IsPublic {
				temp = append(temp, note)
			}
		}

		res = append(res, response{Journey: j, Route: points, Travelers: travelers, Notes: temp})
	}

	return c.JSON(res)
}
