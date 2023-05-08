package memorystorage

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/internal/storage"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()
	memStorage := NewStorage()
	t.Run("simple case", func(t *testing.T) {
		event := entity.Event{
			ID:          uuid.MustParse("8d7b698a-df75-11ed-b5ea-0242ac120002"),
			Title:       "Test",
			Description: "Test",
			StartAt:     time.Now(),
			FinishAt:    time.Now(),
			UserID:      uuid.UUID{},
		}
		err := memStorage.Create(ctx, event)
		require.NoError(t, err)
		require.Equal(t, event, memStorage.mp[uuid.MustParse("8d7b698a-df75-11ed-b5ea-0242ac120002")])

		events, err := memStorage.GetOnDay(ctx, time.Now())
		require.NoError(t, err)
		require.Equal(t, events[0], event)
		err = memStorage.Delete(ctx, uuid.MustParse("8d7b698a-df75-11ed-b5ea-0242ac120002"))
		require.NoError(t, err)
		require.Equal(t, entity.Event{}, memStorage.mp[uuid.MustParse("8d7b698a-df75-11ed-b5ea-0242ac120002")])
	})
	events := []entity.Event{
		{
			ID:       uuid.New(),
			Title:    "+",
			StartAt:  time.Date(2023, time.May, 12, 13, 0, 0, 0, time.UTC),
			FinishAt: time.Time{},
			UserID:   uuid.UUID{},
		},
		{
			ID:       uuid.New(),
			Title:    "+",
			StartAt:  time.Date(2023, time.May, 3, 13, 0, 0, 0, time.UTC),
			FinishAt: time.Time{},
			UserID:   uuid.UUID{},
		},
		{
			ID:       uuid.New(),
			Title:    "+",
			StartAt:  time.Date(2023, time.May, 8, 13, 0, 0, 0, time.UTC),
			FinishAt: time.Time{},
			UserID:   uuid.UUID{},
		},
		{
			ID:       uuid.New(),
			Title:    "week",
			StartAt:  time.Date(2023, time.April, 6, 13, 0, 0, 0, time.UTC),
			FinishAt: time.Time{},
			UserID:   uuid.UUID{},
		},
		{
			ID:       uuid.New(),
			Title:    "week",
			StartAt:  time.Date(2023, time.April, 9, 23, 0, 0, 0, time.UTC),
			FinishAt: time.Time{},
			UserID:   uuid.UUID{},
		},
	}
	for _, event := range events {
		_ = memStorage.Create(ctx, event)
	}
	t.Run("Get on month", func(t *testing.T) {
		months, err := memStorage.GetOnMonth(ctx, time.Date(2023, time.May, 2, 0, 0, 0, 0, time.UTC))
		require.NoError(t, err)
		require.Equal(t, 3, len(months))
		for i, month := range months {
			require.Equal(t, events[i].Title, month.Title)
		}
	})
	t.Run("Get on week", func(t *testing.T) {
		months, err := memStorage.GetOnWeek(ctx, time.Date(2023, time.April, 5, 0, 0, 0, 0, time.UTC))
		require.NoError(t, err)
		require.Equal(t, 2, len(months))
		for i, month := range months {
			require.Equal(t, events[len(events)-i-1].Title, month.Title)
		}
	})
	t.Run("No events on day", func(t *testing.T) {
		_, err := memStorage.GetOnDay(ctx, time.Date(2024, time.December, 5, 1, 21, 1, 0, time.UTC))
		require.Equal(t, err, storage.ErrNothingEventsForDay)
	})

	t.Run("concurrent test", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(len(events))
		for i := 0; i < len(events); i++ {
			go func(i int) {
				defer wg.Done()
				err := memStorage.Delete(ctx, events[i].ID)
				require.NoError(t, err)
			}(i)
		}
		wg.Wait()
	})
}
