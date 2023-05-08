package internalhttp

import (
	"context"
	"github.com/google/uuid"
	"net"
	"net/http"
	"time"

	"github.com/vsPEach/Kursovaya/config"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/internal/server/http/handlers"
)

type Storage interface {
	Create(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, id uuid.UUID, event entity.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error)
}

type Logger interface {
	Error(...interface{})
	Info(...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type HTTPServer struct {
	logger Logger
	server http.Server
}

func NewHTTPServer(logger Logger, conf config.ServerConf, storage Storage) *HTTPServer {
	h := handlers.NewHTTPHandlers(logger, storage)
	return &HTTPServer{
		logger: logger,
		server: http.Server{
			Addr:              net.JoinHostPort(conf.Host, conf.Port),
			Handler:           loggingMiddleware(h.Routes(), logger),
			ReadHeaderTimeout: time.Second * 2,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
