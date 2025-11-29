package main

import (
	"fmt"
	"log/slog"
	"os"

	mwsclient "github.com/Util787/mws-content-registry/internal/adapters/http-clients/mws-client"
	"github.com/Util787/mws-content-registry/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Конфиг MWS

	cfg := config.MustLoadConfig()

	// Создание клиента
	mws := mwsclient.NewMWSClient(logger, cfg.HTTPClientsConfig)

	// GET-запрос на первую страницу таблицы
	res, err := mws.TakeRecords(
		"viwGe6CA09LxA", // viewId
		1,               // pageNum
		100,             // pageSize
		nil,             // сортировка
		nil,             // конкретные recordIds
		nil,             // поля
	)
	if err != nil {
		fmt.Println("Ошибка при запросе MWS:", err)
		return
	}

	// Вывод записей
	logger.Debug("Любая хуйня тест что угодно", slog.Any("Любая хуйня", res))
	for i, rec := range res.Data.Records {
		fmt.Printf("[%d] ID=%d URL=%s Author=%s\n", i, rec.Fields.ID, rec.Fields.URL, rec.Fields.Author)
	}
}
