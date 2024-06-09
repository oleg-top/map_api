package models

import "github.com/google/uuid"

type Traveler struct {
	ID          uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	ChatID      int       `json:"chat_id"`
	Name        string    `json:"name"`
	JourneyName string    `json:"journey_name"`
}
