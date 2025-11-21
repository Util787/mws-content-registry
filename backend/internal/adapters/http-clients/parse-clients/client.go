package parseclients

import (
	"log/slog"

	httpclients "github.com/Util787/mws-content-registry/internal/adapters/http-clients"
	"github.com/Util787/mws-content-registry/internal/config"
	"github.com/go-resty/resty/v2"
)

type ParseClient struct {
	log    *slog.Logger
	client *resty.Client
	//urls to parse...
}

func NewParseClient(log *slog.Logger, cfg config.HTTPClientsConfig) *ParseClient {

	rclient := httpclients.NewRestyClient()

	return &ParseClient{
		log:    log,
		client: rclient,
		//urls...
	}
}
