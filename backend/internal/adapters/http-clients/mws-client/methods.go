package mwsclient

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/Util787/mws-content-registry/internal/models"
)

// TakeAll получает все записи из MWS API
func (mwsClient *MWSClient) TakeAll() (models.Response, error) {
	// GET-запрос
	res, err := mwsClient.client.R().Get(mwsClient.MWSUrl)
	if err != nil {
		mwsClient.log.Error("Failed to make request to MWS API", "error", err)
		return models.Response{}, err
	}

	// Парсинг JSON
	var response models.Response
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		mwsClient.log.Error("Failed to unmarshal MWS response", "error", err)
		return models.Response{}, err
	}

	// Логирование базовой инфы
	mwsClient.log.Info("Fetched MWS records",
		"total_records", len(response.Data.Records),
		"code", response.Code,
		"success", response.Success,
	)

	// Дополнительный вывод первых 3-х строк для отладки
	for i, rec := range response.Data.Records {
		if i >= 3 {
			break
		}
		fmt.Println("Record:", i)
		fmt.Println("RecordID:", rec.RecordID)
		fmt.Println("Fields:", rec.Fields)
		fmt.Println("----------------------------------------------------")
	}

	return response, nil
}

// TakeByID Пример, как можно будет достать запись по ID
func (mwsClient *MWSClient) TakeByID(id string) (models.Record, error) {
	// Добавляем query-параметр ?id=<id>
	url := fmt.Sprintf("%s?id=%s", mwsClient.MWSUrl, id)
	res, err := mwsClient.client.R().Get(url)
	if err != nil {
		mwsClient.log.Error("Failed to fetch MWS record by ID", "id", id, "error", err)
		return models.Record{}, err
	}

	var response models.Response
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		mwsClient.log.Error("Failed to unmarshal MWS response for ID", "id", id, "error", err)
		return models.Record{}, err
	}

	if len(response.Data.Records) == 0 {
		return models.Record{}, fmt.Errorf("record with ID %s not found", id)
	}

	return response.Data.Records[0], nil
}
