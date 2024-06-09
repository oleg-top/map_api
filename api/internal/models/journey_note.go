package models

import "github.com/google/uuid"

type JourneyNote struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Author      int       `json:"author"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	JourneyName string    `json:"journey_name"`
	IsPublic    bool      `json:"is_public"`
}
