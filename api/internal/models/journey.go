package models

import (
	"github.com/google/uuid"
)

type Journey struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
