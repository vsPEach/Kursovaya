package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notice struct {
	ID     uuid.UUID
	Title  string
	Date   time.Time
	UserID uuid.UUID
}
