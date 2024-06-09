package pointHandler

import (
	"encoding/json"
	"time"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type AddPointRequest struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Location    string `json:"location"`
	Description string `json:"description"`
	JourneyName string `json:"journey_name"`
}

func AddPoint(c fiber.Ctx) error {
	var req AddPointRequest

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

	var journey models.Journey
	result := database.DB.Where("name = ?", req.JourneyName).First(&journey)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Путешествие с таким названием не найдено",
		})
	}

	point := models.Point{
		StartDate:   startDate,
		EndDate:     endDate,
		Location:    req.Location,
		Description: req.Description,
		JourneyName: req.JourneyName,
	}

	database.DB.Create(&point)
	return c.Status(201).JSON(fiber.Map{
		"message": "Точка маршрута успешно создана",
	})
}
