package memorystorage

import (
	"context"

	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/internal/storage"
	"github.com/vsPEach/Kursovaya/pkg/timeconv"
)

type Storage struct {
	mu sync.RWMutex
	mp map[uuid.UUID]entity.Event
}

func (s *Storage) Create(ctx context.Context, event entity.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.mp[event.ID]
	if ok {
		return storage.ErrEventAlreadyExist
	}
	s.mp[event.ID] = event
	return nil
}

func (s *Storage) Update(ctx context.Context, id uuid.UUID, event entity.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.mp[id]
	if !ok {
		return storage.ErrEventDoesNotExist
	}
	s.mp[id] = event
	return nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.mp[id]
	if !ok {
		return storage.ErrEventDoesNotExist
	}
	delete(s.mp, id)
	return nil
}

func (s *Storage) GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, event := range s.mp {
		if timeconv.ToDateOnly(event.StartAt) == timeconv.ToDateOnly(date) {
			events = append(events, event)
		}
	}
	if len(events) == 0 {
		return nil, storage.ErrNothingEventsForDay
	}
	return events, nil
}

func (s *Storage) GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	s.mu.RLock()
	defer s.mu.RUnlock()
	weekDate := date.AddDate(0, 0, 7)
	for _, event := range s.mp {
		if event.StartAt.After(date) && event.StartAt.Before(weekDate) {
			events = append(events, event)
		}
	}
	if len(events) == 0 {
		return nil, storage.ErrNothingEventsForWeek
	}
	return events, nil
}

func (s *Storage) GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	s.mu.RLock()
	defer s.mu.RUnlock()

	monthDate := date.AddDate(0, 1, 0)
	for _, event := range s.mp {
		if event.StartAt.After(date) && event.StartAt.Before(monthDate) {
			events = append(events, event)
		}
	}
	if len(events) == 0 {
		return nil, storage.ErrNothingEventsForMonth
	}
	return events, nil
}

func NewStorage() Storage {
	return Storage{
		mp: make(map[uuid.UUID]entity.Event),
	}
}
