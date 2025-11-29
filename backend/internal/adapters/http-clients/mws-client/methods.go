package mwsclient

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/Util787/mws-content-registry/internal/common"
	"github.com/Util787/mws-content-registry/internal/models"
)

// TakeRecords получает записи из MWS Tables с поддержкой пагинации, сортировки и фильтров
func (mwsClient *MWSClient) TakeRecords(
	pageNum int,
	pageSize int,
	sort []map[string]string,
	recordIds []string,
	fields []string,
) ([]models.MWSTableRecord, error) {
	log := mwsClient.log.With("op", common.GetOperationName())

	req := mwsClient.client.R()

	req.SetQueryParam("viewId", mwsClient.MWSViewID)

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

	log.Debug("req", slog.Any("body", req.Body), slog.Any("headers", req.Header))

	res, err := req.Get(mwsClient.MWSUrl)
	if err != nil {
		return nil, err
	}

	var response models.MWSTableResponse
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		return nil, err
	}

	return response.Data.Records, nil
}

// AddRecords добавляет новые записи в таблицу
func (mwsClient *MWSClient) AddRecords(records []models.MWSTableNewRecord) error {
	log := mwsClient.log.With("op", common.GetOperationName())

	reqBody := map[string]interface{}{
		"viewId":  mwsClient.MWSViewID,
		"records": records,
	}

	res, err := mwsClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post(mwsClient.MWSUrl)
	if err != nil {
		return err
	}

	var respStruct models.MWSTableResponse

	json.Unmarshal(res.Body(), &respStruct)

	if res.Error() != nil {
		log.Error("mws-api response error", slog.Int("status", res.StatusCode()), slog.Any("body", respStruct))
		return err
	}

	mwsClient.log.Debug("mws-table-response", slog.Any("resp", respStruct))
	var response models.MWSTableResponse
	if err := json.Unmarshal(res.Body(), &response); err != nil {
		return err
	}

	return nil
}
