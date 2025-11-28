package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Util787/mws-content-registry/backend/internal/config"

	mwsclient "github.com/Util787/mws-content-registry/backend/internal/adapters/http-clients/mws-client"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	//cfg := config.MustLoadConfig()

	mws := mwsclient.NewMWSClient(logger, config.HTTPClientsConfig{MWSUrl: "https://tables.mws.ru/fusion/v1/datasheets/dstAAM6Vof8yCssdVr/records", MWSToken: "uskIIAZwODC7ElSFqXOVQJs"})

	res, err := mws.TakeAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
