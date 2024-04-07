package app

import (
	"PrettyLoger/internal/services"
	"io"
	"log"
	"log/slog"
)

func InitLogger(
	LevelDebug slog.Leveler,
	OutputWriter io.Writer,
) *slog.Logger {
	return slog.New(&services.PrettyLogger{
		Handler: slog.NewJSONHandler(OutputWriter, &slog.HandlerOptions{Level: LevelDebug}),
		Logger:  log.New(OutputWriter, "", 0),
	})
}
