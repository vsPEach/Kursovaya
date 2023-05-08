package internalgrpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/vsPEach/Kursovaya/config"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/internal/server"
	"github.com/vsPEach/Kursovaya/internal/server/grpc/pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate protoc -I ../../../api EventService.proto --go_out=./pb --go-grpc_out=./pb

type Storage interface {
	Create(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, id uuid.UUID, event entity.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error)
}

type Logger interface {
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Debug(...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type CalendarService struct {
	*pb.UnimplementedCalendarServiceServer
	storage Storage
	logger  Logger
}

func (c *CalendarService) CreateEvent(ctx context.Context, request *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	event, err := c.toEntityEvent(request.Event)
	if err != nil {
		return new(pb.CreateEventResponse), err
	}
	if err := c.storage.Create(ctx, event); err != nil {
		return new(pb.CreateEventResponse), err
	}
	return &pb.CreateEventResponse{Event: request.Event}, nil
}

func (c *CalendarService) UpdateEvent(ctx context.Context, request *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	event, err := c.toEntityEvent(request.Event)
	if err != nil {
		return new(pb.UpdateEventResponse), err
	}
	id, err := c.toUUID(request.Event.Id)
	if err != nil {
		return new(pb.UpdateEventResponse), err
	}
	if err := c.storage.Update(ctx, id, event); err != nil {
		return new(pb.UpdateEventResponse), err
	}
	return &pb.UpdateEventResponse{Event: request.Event}, nil
}

func (c *CalendarService) DeleteEvent(ctx context.Context, request *pb.DeleteEventRequest) (*emptypb.Empty, error) {
	id, err := c.toUUID(request.Id)
	if err != nil {
		return new(emptypb.Empty), err
	}
	if err := c.storage.Delete(ctx, id); err != nil {
		return new(emptypb.Empty), err
	}
	return &emptypb.Empty{}, nil
}

func (c *CalendarService) GetDayEvents(ctx context.Context, request *pb.GetDayEventsRequest) (*pb.GetDayEventsResponse, error) {
	events, err := c.storage.GetOnDay(ctx, request.Time.AsTime())
	if err != nil {
		return new(pb.GetDayEventsResponse), err
	}
	return &pb.GetDayEventsResponse{Event: c.toPbEvents(events)}, nil
}

func (c *CalendarService) GetWeekEvents(ctx context.Context, request *pb.GetWeekEventsRequest) (*pb.GetWeekEventsResponse, error) {
	events, err := c.storage.GetOnWeek(ctx, request.Time.AsTime())
	if err != nil {
		return new(pb.GetWeekEventsResponse), err
	}
	return &pb.GetWeekEventsResponse{Event: c.toPbEvents(events)}, nil
}

func (c *CalendarService) GetMonthEvents(ctx context.Context, request *pb.GetMonthEventsRequest) (*pb.GetMonthEventsResponse, error) {
	events, err := c.storage.GetOnMonth(ctx, request.Time.AsTime())
	if err != nil {
		return new(pb.GetMonthEventsResponse), err
	}
	return &pb.GetMonthEventsResponse{Event: c.toPbEvents(events)}, nil
}

func NewCalendarService(conf config.RPC, logger Logger, storage Storage) error {
	lsn, err := net.Listen("tcp", net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)))
	if err != nil {
		return err
	}
	rpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		Logging(logger),
	))
	pb.RegisterCalendarServiceServer(rpcServer, &CalendarService{
		logger:  logger,
		storage: storage,
	})
	if err = rpcServer.Serve(lsn); err != nil {
		return err
	}
	return nil
}

func (c *CalendarService) Shutdown(ctx context.Context) {
	//TODO: implement me
}

func (c *CalendarService) toUUID(id string) (uuid.UUID, error) {
	UUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, server.ErrCantParseUUID
	}
	return UUID, nil
}

func (c *CalendarService) toPbEvents(events entity.Events) []*pb.Event {
	pbEvents := make([]*pb.Event, 0, 10)
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.Event{
			Id:          event.ID.String(),
			Title:       event.Title,
			Description: event.Title,
			StartAt:     timestamppb.New(event.StartAt),
			FinishAt:    timestamppb.New(event.FinishAt),
			UserId:      event.UserID.String(),
		})
	}
	return pbEvents
}

func (c *CalendarService) toEntityEvent(event *pb.Event) (entity.Event, error) {
	id, err := c.toUUID(event.Id)
	if err != nil {
		return entity.Event{}, err
	}
	userID, err := c.toUUID(event.UserId)
	if err != nil {
		return entity.Event{}, err
	}
	return entity.Event{
		ID:          id,
		Title:       event.Title,
		Description: event.Description,
		StartAt:     event.StartAt.AsTime(),
		FinishAt:    event.FinishAt.AsTime(),
		UserID:      userID,
	}, nil
}
