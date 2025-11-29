package usecase

import "github.com/Util787/mws-content-registry/internal/models"

type MWSTablesUsecase struct {
	MWSTablesClient
	YouTubeParseClient
	LLMClient
}

type MWSTablesClient interface {
	AddRecords(records []models.MWSTableNewRecord) error
	TakeRecords(pageNum int, pageSize int, sort []map[string]string, recordId string, fields []string) ([]models.MWSTableRecord, error)
	UpdateRecords(records []models.MWSTableUpdateRecord) error
}

type YouTubeParseClient interface {
	ScrabVideosWithComments() ([]models.YTVideosWithComments, error)
	ScrabVideoByURL(videoURL string) (*models.YTVideosWithComments, error)
}

type LLMClient interface {
	GenerateContentAnalyze(rec models.MWSTableRecord) (models.AnalyzeData, error)
}

func NewMWSTablesUsecase(mwsClient MWSTablesClient, ytClient YouTubeParseClient, llmClient LLMClient) *MWSTablesUsecase {
	return &MWSTablesUsecase{
		MWSTablesClient:    mwsClient,
		YouTubeParseClient: ytClient,
		LLMClient:          llmClient,
	}
}
