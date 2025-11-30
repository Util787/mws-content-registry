package usecase

import (
	"log/slog"

	"github.com/Util787/mws-content-registry/internal/models"
)

type MWSTablesUsecase struct {
	MWSTablesClient
	YouTubeParseClient
	LLMClient
}

type MWSTablesClient interface {
	AddRecords(records []models.MWSTableNewRecord) error
	TakeRecords(pageNum int, pageSize int, sort map[string]string, recordId string, fields []string) ([]models.MWSTableRecord, error)
	UpdateRecords(records []models.MWSTableUpdateRecord) error
}

type YouTubeParseClient interface {
	ScrabVideosWithComments() ([]models.YTVideosWithComments, error)
	ScrabVideoByURL(videoURL string) (*models.YTVideosWithComments, error)
}

type LLMClient interface {
	GenerateContentAnalyze(rec models.MWSTableRecord) (models.AnalyzeData, error)
	GenerateChatAnswer(recs []models.MWSTableRecord, userMessage string) (string, error)
}

func NewMWSTablesUsecase(mwsClient MWSTablesClient, ytClient YouTubeParseClient, llmClient LLMClient) *MWSTablesUsecase {
	return &MWSTablesUsecase{
		MWSTablesClient:    mwsClient,
		YouTubeParseClient: ytClient,
		LLMClient:          llmClient,
	}
}

type AiChatUsecase struct {
	MessageStorage
	LLMClient
	MWSTablesClient
	logger      *slog.Logger
	queueToSave chan models.Message
}

type MessageStorage interface {
	SaveMessage(mes models.Message) error
	GetChatHistory(chat int64) ([]models.Message, error)
}

func NewAiChatUsecase(ms MessageStorage, mwsClient MWSTablesClient, llmClient LLMClient, queueSize int, logger *slog.Logger) *AiChatUsecase {
	au := &AiChatUsecase{
		MessageStorage:  ms,
		MWSTablesClient: mwsClient,
		LLMClient:       llmClient,
		queueToSave:     make(chan models.Message, queueSize),
		logger:          logger,
	}

	au.startMessageSaver()

	return au
}
