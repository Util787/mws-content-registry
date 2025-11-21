package mwsclient

import (
	"log/slog"

	httpclients "github.com/Util787/mws-content-registry/internal/adapters/http-clients"
	"github.com/Util787/mws-content-registry/internal/config"
	"github.com/go-resty/resty/v2"
)

type MWSClient struct {
	log    *slog.Logger
	client *resty.Client
	MWSUrl string
}

func NewMWSClient(log *slog.Logger, cfg config.HTTPClientsConfig) *MWSClient {

	rclient := httpclients.NewRestyClient()

	return &MWSClient{
		log:    log,
		client: rclient,
		MWSUrl: cfg.MWSUrl,
	}
}
