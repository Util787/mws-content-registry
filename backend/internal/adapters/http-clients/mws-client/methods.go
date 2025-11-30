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
	sort map[string]string,
	recordId string,
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
		req.SetQueryParams(sort)
	}
	if recordId != "" {
		req.SetQueryParam("recordIds", recordId)
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

	if len(response.Data.Records) == 0 {
		log.Warn("No records found in MWS response", slog.Int("status", res.StatusCode()), slog.Any("body", response))
		return []models.MWSTableRecord{}, fmt.Errorf("records not found")
	}

	// log.Debug("got records", slog.Any("records", response.Data.Records))

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

func (mwsClient *MWSClient) UpdateRecords(records []models.MWSTableUpdateRecord) error {
	log := mwsClient.log.With("op", common.GetOperationName())

	reqBody := map[string]interface{}{
		"viewId":  mwsClient.MWSViewID,
		"records": records,
	}

	rawResp, err := mwsClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Patch(mwsClient.MWSUrl)
	if err != nil {
		return err
	}

	log.Debug("updated records", slog.Any("records", records))

	var respStruct models.MWSTableResponse

	err = json.Unmarshal(rawResp.Body(), &respStruct)
	if err != nil {
		log.Error("failed to unmarshal resp", slog.String("error", err.Error()))
		return err
	}

	if rawResp.Error() != nil {
		log.Error("mws-api response error", slog.Int("status", rawResp.StatusCode()), slog.Any("body", respStruct))
		return err
	}

	mwsClient.log.Debug("mws-table-response", slog.Any("resp", respStruct))
	var response models.MWSTableResponse
	if err := json.Unmarshal(rawResp.Body(), &response); err != nil {
		return err
	}

	return nil
}
