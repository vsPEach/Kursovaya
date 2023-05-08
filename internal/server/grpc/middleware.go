package internalgrpc

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func Logging(logger Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not grab metadata from context")
		}

		start := time.Now()
		resp, err = handler(ctx, req)
		s, _ := status.FromError(err)
		if err != nil {
			return nil, err
		}
		latency := time.Since(start).Seconds()
		logger.Infow("gRPC request",
			zap.Strings("user-agent", meta.Get("user-agent")),
			zap.Strings("content-type", meta.Get("content-type")),
			zap.Float64("latency", latency),
			zap.String("Status", s.Code().String()),
		)
		return resp, err
	}
}
