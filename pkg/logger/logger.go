package logger

import (
	"log/slog"
	"os"
)

func NewLogger(serviceName string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug}).WithAttrs([]slog.Attr{
		slog.String("service", serviceName),
	}))
	return logger
}
