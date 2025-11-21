package main

import (
	"log/slog"
	"os"

	"github.com/Util787/mws-content-registry/internal/config"
)

func main() {
	logger := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	cfg := config.MustLoadConfig()
}
