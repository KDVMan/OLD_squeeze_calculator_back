package core_services_logger

import (
	"log/slog"
	"os"
)

const (
	envLocal      = "local"
	envProduction = "production"
)

func New(env string) *slog.Logger {
	var logger *slog.Logger

	if env == envLocal {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else if env == envProduction {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
