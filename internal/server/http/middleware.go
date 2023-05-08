package internalhttp

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type StatusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func loggingMiddleware(next http.Handler, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		t := time.Now()
		next.ServeHTTP(recorder, r)
		latency := time.Since(t).Seconds()
		logger.Infow("HTTP request: ",
			zap.String("IP", r.RemoteAddr),
			zap.String("Method", r.Method),
			zap.String("Path", r.RequestURI),
			zap.String("Version", r.Proto),
			zap.Int("Status", recorder.status),
			zap.Float64("Latency", latency),
			zap.String("User-Agent", r.UserAgent()),
		)
	})
}
