package llmclient

import (
	"log/slog"

	httpclients "github.com/Util787/mws-content-registry/internal/adapters/http-clients"
	"github.com/Util787/mws-content-registry/internal/config"
	"github.com/go-resty/resty/v2"
)

type LLMClient struct {
	log      *slog.Logger
	client   *resty.Client
	LLMUrl   string
	LLMModel string
}

func NewLLMClient(log *slog.Logger, cfg config.HTTPClientsConfig) *LLMClient {

	rclient := httpclients.NewRestyClient()

	rclient = rclient.SetHeader("Authorization", "Bearer "+cfg.LLMClient.LLMApiKey)

	return &LLMClient{
		log:      log,
		client:   rclient,
		LLMUrl:   cfg.LLMClient.LLMUrl,
		LLMModel: cfg.LLMClient.LLMModel,
	}
}
