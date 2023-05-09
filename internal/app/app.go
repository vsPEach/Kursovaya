package app

//nolint:gocritic

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vsPEach/Kursovaya/config"
	internalgrpc "github.com/vsPEach/Kursovaya/internal/server/grpc"
	httpserver "github.com/vsPEach/Kursovaya/internal/server/http"
	sqlstorage "github.com/vsPEach/Kursovaya/internal/storage/sql"
	"github.com/vsPEach/Kursovaya/pkg/logger"
)

type Logger interface {
	Info(...interface{})
	Error(...interface{})
}

func Run(config config.Config) {
	logg := logger.New(config.Logger)
	wg := sync.WaitGroup{}
	wg.Add(4)

	storage, err := sqlstorage.New(config.Database, logg)
	if err != nil {
		logg.Error(err.Error())
	}

	httpServer := httpserver.NewHTTPServer(logg, config.Server, storage)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer wg.Done()
		defer cancel()

		if err := storage.Close(); err != nil {
			logg.Errorw("failed to stop storage: " + err.Error())
			cancel()
		}

		logg.Info("Stop Storage")

		if err := httpServer.Stop(ctx); err != nil {
			logg.Errorw("failed to stop http httpServer: " + err.Error())
			cancel()
		}
		logg.Info("Stop Http server")

	}()
	logg.Info("calendar is running...")

	go func() {
		defer wg.Done()
		if err := storage.Connect(ctx); err != nil {
			logg.Error("Can't ping database: " + err.Error())
			cancel()
			return
		}
		logg.Info("Connected to database")
	}()

	go func() {
		defer wg.Done()
		logg.Info("http server start")
		if err := httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			return
		}
	}()

	go func() {
		defer wg.Done()
		logg.Info("gRPC server start")
		if err := internalgrpc.NewCalendarService(config.Rpc, logg, storage); err != nil {
			logg.Error("failed to start gRPC server: " + err.Error())
			cancel()
			return
		}
	}()

	wg.Wait()
}
