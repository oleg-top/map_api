package profileHandler

import (
	"encoding/json"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

// TODO: проверка на валидность города и валидность данных при реге
func UpdateProfile(c fiber.Ctx) error {
	var req models.Chat

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	var chat models.Chat
	result := database.DB.Model(&chat).Where("id = ?", req.ID).First(&chat)

	if result.Error != nil {
		database.DB.Create(&req)

		return c.JSON(fiber.Map{
			"message": "Пользователь был успешно создан",
		})
	}

	database.DB.Model(&chat).Updates(req)

	return c.JSON(fiber.Map{
		"message": "Данные пользователя были успешно обновлены",
	})
}
