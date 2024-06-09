package models

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	UserChatID  int
	OweChatID   int
	Amount      int
	Date        time.Time
	JourneyName string
	IsPaid      bool
}
