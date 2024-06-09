package pointHandler

import (
	"encoding/json"
	"time"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type RemovePointRequest struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Location    string `json:"location"`
	JourneyName string `json:"journey_name"`
}

func RemovePoint(c fiber.Ctx) error {
	var req RemovePointRequest

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	startDate, err := time.Parse("02.01.2006", req.StartDate)
	if err != nil {
		return err
	}

	endDate, err := time.Parse("02.01.2006", req.EndDate)
	if err != nil {
		return err
	}

	database.DB.
		Where("start_date = ? AND end_date = ? AND location = ? AND journey_name = ?", startDate, endDate, req.Location, req.JourneyName).
		Delete(&models.Point{})

	return c.JSON(fiber.Map{
		"message": "Точка маршрута была успешно удалена",
	})
}
