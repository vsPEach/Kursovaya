package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgresql driver
	"github.com/vsPEach/Kursovaya/config"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/pkg/timeconv"
)

type Logger interface {
	Info(...interface{})
}

type Storage struct {
	logg Logger
	conn *sqlx.DB
}

func (s *Storage) GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	err := s.conn.SelectContext(ctx, &events,
		"select * from events where TO_CHAR(start_at, 'DD-MM-YYYY') = $1",
		timeconv.ToDateOnly(date))
	return events, err
}

func (s *Storage) GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	err := s.conn.SelectContext(ctx, &events,
		"select * from events where TO_CHAR(start_at, 'DD-MM-YYYY') between $1 and $2",
		timeconv.ToDateOnly(date),
		timeconv.GetWeek(date))
	return events, err
}

func (s *Storage) GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error) {
	var events entity.Events
	err := s.conn.SelectContext(ctx, &events,
		"select * from events where TO_CHAR(start_at, 'DD-MM-YYYY') between $1 and $2",
		timeconv.ToDateOnly(date),
		timeconv.GetMonth(date))
	return events, err
}

func (s *Storage) Create(ctx context.Context, event entity.Event) error {
	query := `insert into events 
    	(id, title, description, start_at, finish_at, user_id) 
	values
	    (:id, :title, :description, :start_at, :finish_at,:user_id)`
	if _, err := s.conn.NamedExecContext(ctx, query, event); err != nil {
		return err
	}
	s.logg.Info("user successfully created")
	return nil
}

func (s *Storage) Update(ctx context.Context, id uuid.UUID, event entity.Event) error {
	query := fmt.Sprintf(`update events 
								set title=:title, description=:description, start_at=:start_at, finish_at=:finish_at
								where id='%d'`, id)
	if _, err := s.conn.NamedExecContext(ctx, query, event); err != nil {
		return err
	}
	s.logg.Info("user successfully updated")
	return nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	query := fmt.Sprintf("delete from events where id='%s'", id)
	if _, err := s.conn.ExecContext(ctx, query); err != nil {
		return err
	}
	s.logg.Info("user successfully delete")
	return nil
}

func New(conf config.DatabaseConf, logger Logger) (*Storage, error) {
	conn, err := sqlx.Open(conf.Type, conf.Url)
	if err != nil {
		return nil, err
	}
	return &Storage{
		conn: conn,
		logg: logger,
	}, nil
}

func (s *Storage) Connect(ctx context.Context) error {
	err := s.conn.PingContext(ctx)
	return err
}

func (s *Storage) Close() error {
	return s.conn.Close()
}
