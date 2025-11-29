package usecase

import "github.com/Util787/mws-content-registry/internal/models"

type ParseUsecase struct {
	YouTubeParseClient
}

type YouTubeParseClient interface {
	ScrabVideosWithComments() ([]models.YTVideosWithComments, error)
}
