package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/mmaxemm/internship_tasks/interfaces"
)
func main() {
	jsonFile, err := os.OpenFile("logs.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	jsonHandler := slog.NewJSONHandler(jsonFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	compositeHandler := interfaces.NewCompositeHandler(textHandler, jsonHandler)

	logger := slog.New(compositeHandler)

	logger.Info("Program started",
		slog.String("version", "1.0.0"),
		slog.Int("port", 8080),
	)

	logger.Debug("Debug info",
		slog.String("module", "database"),
	)

	logger.Warn("Warning",
		slog.String("action", "cache_miss"),
	)

	logger.Error("Error",
		slog.String("operation", "user_create"),
		slog.String("reason", "invalid_email"),
	)

	ctx := context.Background()
	logger.InfoContext(ctx, "Quitting",
		slog.String("status", "success"),
	)
}
