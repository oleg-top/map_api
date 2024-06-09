package database

import (
	"log/slog"
	"os"
	"time"

	"github.com/Central-University-IT-prod/backend-oleg-top/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	pgUrl := os.Getenv("POSTGRES_CONN")
	if pgUrl == "" {
		slog.Error("missed POSTGRES_CONN env")
		os.Exit(1)
	}

	time.Sleep(5 * time.Second)

	var err error

	DB, err = gorm.Open(postgres.Open(pgUrl), &gorm.Config{})
	if err != nil {
		slog.Error("error while opening db")
		os.Exit(1)
	}

	DB.AutoMigrate(&models.Chat{}, &models.Point{}, &models.Journey{},
		&models.Traveler{}, &models.JourneyNote{}, &models.Debt{})
}
