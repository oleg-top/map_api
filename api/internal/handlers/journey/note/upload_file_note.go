package noteHandler

import (
	"fmt"
	"time"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/database"
	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"github.com/gofiber/fiber/v3"
)

func UploadFileNote(c fiber.Ctx) error {
	noteID := c.Params("note_id")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	newFileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), file.Filename)
	destination := fmt.Sprintf("./uploads/%s", newFileName)

	if err = c.SaveFile(file, destination); err != nil {
		return err
	}

	var note models.JourneyNote
	database.DB.Where("id = ?", noteID).First(&note)

	note.Path = newFileName

	database.DB.Save(&note)

	return c.JSON(fiber.Map{
		"message": "Файл был успешно загружен",
	})
}
