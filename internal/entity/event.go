package entity

import (
	"time"

	"github.com/google/uuid"
)

type Events []Event

type Event struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	StartAt     time.Time `db:"start_at" json:"start_at"`
	FinishAt    time.Time `db:"finish_at" json:"finish_at"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
}
