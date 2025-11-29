package parseclients

import (
	"context"
	"log/slog"

	"github.com/Util787/mws-content-registry/internal/config"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeParseClient struct {
	log           *slog.Logger
	ytService     *youtube.Service
	videosLimit   int64
	commentsLimit int64
	chart         string
	regionCode    string
}

func NewYouTubeParseClient(ctx context.Context, log *slog.Logger, cfg config.HTTPClientsConfig) *YouTubeParseClient {

	yService, err := youtube.NewService(ctx, option.WithAPIKey(cfg.YouTubeParseClient.YouTubeAPIKey))
	if err != nil {
		log.Error("Failed to create YouTube service in parse client", "error", err)
	}

	return &YouTubeParseClient{
		log:           log,
		ytService:     yService,
		videosLimit:   cfg.YouTubeParseClient.VideosLimit,
		commentsLimit: cfg.YouTubeParseClient.CommentsLimit,
		chart:         cfg.YouTubeParseClient.Chart,
		regionCode:    cfg.YouTubeParseClient.RegionCode,
	}
}
