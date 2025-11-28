package mwsclient

import (
	"fmt"
	"log/slog"

	httpclients "github.com/Util787/mws-content-registry/backend/internal/adapters/http-clients"
	"github.com/Util787/mws-content-registry/backend/internal/config"
	"github.com/go-resty/resty/v2"
)

type MWSClient struct {
	log    *slog.Logger
	client *resty.Client
	MWSUrl string
}

func NewMWSClient(log *slog.Logger, cfg config.HTTPClientsConfig) *MWSClient {

	rclient := httpclients.NewRestyClient()
	rclient.SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.MWSToken))

	return &MWSClient{
		log:    log,
		client: rclient,
		MWSUrl: cfg.MWSUrl,
	}
}
