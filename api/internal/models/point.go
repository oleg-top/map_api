package models

import (
	"time"

	"github.com/google/uuid"
)

type Point struct {
	ID          uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	JourneyName string    `json:"journey_name"`
}
