package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vsPEach/Kursovaya/internal/entity"
)

type Storage interface {
	Create(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, id uuid.UUID, event entity.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error)
}

type Sender struct {
	storage Storage
	logger  Logger
}

func NewSender(storage Storage, logger Logger) *Sender {
	return &Sender{storage: storage, logger: logger}
}

func (s *Sender) SendNotify() {

}
