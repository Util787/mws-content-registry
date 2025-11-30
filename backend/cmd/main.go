package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	llmclient "github.com/Util787/mws-content-registry/internal/adapters/http-clients/llm-client"
	mwsclient "github.com/Util787/mws-content-registry/internal/adapters/http-clients/mws-client"
	parseclients "github.com/Util787/mws-content-registry/internal/adapters/http-clients/parse-clients"
	"github.com/Util787/mws-content-registry/internal/adapters/postgresql"
	"github.com/Util787/mws-content-registry/internal/adapters/rest"
	"github.com/Util787/mws-content-registry/internal/config"
	"github.com/Util787/mws-content-registry/internal/usecase"
)

const defaultQueueSize = 100

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg := config.MustLoadConfig()

	// Initialize HTTP clients
	mwsClient := mwsclient.NewMWSClient(logger, cfg.HTTPClientsConfig)
	ytClient := parseclients.NewYouTubeParseClient(context.Background(), logger, cfg.HTTPClientsConfig)
	llmClient := llmclient.NewLLMClient(logger, cfg.HTTPClientsConfig)

	// Initialize postgres
	pgPool, err := postgresql.ConnectPostgreSQL(cfg.PostgresConfig)
	if err != nil {
		panic("Failed to connect to Postgresql: " + err.Error())
	}
	pgStorage := postgresql.NewStorage(pgPool, logger)

	// Initialize usecases
	mwUc := usecase.NewMWSTablesUsecase(mwsClient, ytClient, llmClient)
	aiChatUc := usecase.NewAiChatUsecase(pgStorage, mwsClient, llmClient, defaultQueueSize, logger)

	// Initialize and run REST server
	server := rest.NewRestServer(logger, cfg.HTTPServerConfig, mwUc, aiChatUc)
	go func() {
		logger.Info("HTTP server start", slog.String("host", cfg.HTTPServerConfig.Host), slog.Int("port", cfg.HTTPServerConfig.Port))
		if err := server.Run(); err != nil {
			logger.Error("HTTP server error", slog.String("error", err.Error()))
		}
	}()

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	logger.Info("Shutting down gracefully...")

	logger.Info("Shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("HTTP server shutdown error", slog.String("error", err.Error()))
	}

	logger.Info("Shutdown complete")

}
