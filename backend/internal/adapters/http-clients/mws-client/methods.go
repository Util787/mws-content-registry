package mwsclient

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/Util787/mws-content-registry/internal/models"
)

// NewRecord — структура для добавления новой записи
type NewRecord struct {
	Fields map[string]interface{} `json:"fields"`
}

// TakeRecords получает записи из MWS Tables с поддержкой пагинации, сортировки и фильтров
func (mwsClient *MWSClient) TakeRecords(
	viewId string,
	pageNum int,
	pageSize int,
	sort []map[string]string,
	recordIds []string,
	fields []string,
) (models.Response, error) {

	req := mwsClient.client.R()

	if viewId != "" {
		req.SetQueryParam("viewId", viewId)
	}
	if pageNum > 0 {
		req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageNum))
	}
	if pageSize > 0 {
		req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageSize))
	}
	if len(sort) > 0 {
		sortJSON, _ := json.Marshal(sort)
		req.SetQueryParam("sort", string(sortJSON))
	}
	if len(recordIds) > 0 {
		idsJSON, _ := json.Marshal(recordIds)
		req.SetQueryParam("recordIds", string(idsJSON))
	}
	if len(fields) > 0 {
		fieldsJSON, _ := json.Marshal(fields)
		req.SetQueryParam("fields", string(fieldsJSON))
	}

	res, err := req.Get(mwsClient.MWSUrl)
	if err != nil {
		return models.Response{}, err
	}

	var response models.Response
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		return models.Response{}, err
	}

	return response, nil
}

// AddRecords добавляет новые записи в таблицу
func (mwsClient *MWSClient) AddRecords(viewId string, records []NewRecord) (models.Response, error) {
	reqBody := map[string]interface{}{
		"viewId":  viewId,
		"records": records,
	}

	res, err := mwsClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post(mwsClient.MWSUrl)
	if err != nil {
		return models.Response{}, err
	}

	mwsClient.log.Debug("12e", slog.Any("123", res))
	var response models.Response
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		return models.Response{}, err
	}

	return response, nil
}
