package main

import (
	"PrettyLoger/internal/app"
	"log/slog"
	"os"
)

func main() {
	logger := app.InitLogger(slog.LevelDebug, os.Stdout)
	logger.Info("Test Info")
	logger.Debug("Test Debug")
	logger.Warn("Test Warning")
	logger.Error("Test", slog.String("id", "123"), slog.String("username", "test"))
}
