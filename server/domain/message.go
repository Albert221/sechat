package domain

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id       uuid.UUID `json:"id"`
	SenderId uuid.UUID `json:"senderId"`
	Message  string    `json:"message"`
	SentAt   time.Time `json:"sentAt"`
}
