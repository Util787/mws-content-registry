package llmclient

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/Util787/mws-content-registry/internal/common"
	"github.com/Util787/mws-content-registry/internal/models"
)

type llmResp struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (lc *LLMClient) GenerateContentAnalyze(dataJSON []byte) (models.AnalyzeData, error) {
	log := lc.log.With("op", common.GetOperationName())

	prompt := getAnalyzePrompt(dataJSON)

	resp, err := lc.client.R().SetBody(map[string]interface{}{
		"model": lc.LLMModel,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		}}).
		Post(lc.LLMUrl)

	if err != nil {
		log.Error("Failed to call llm-api for generating content analyze", "error", err)
		return models.AnalyzeData{}, err
	}

	if resp.StatusCode() != 200 {
		log.Error("llm-api returned non-200 status code", slog.Int("status", resp.StatusCode()), slog.Any("body", resp.Body()))
		return models.AnalyzeData{}, err
	}

	var llmResponse llmResp
	err = json.Unmarshal(resp.Body(), &llmResponse)
	if err != nil {
		log.Error("Failed to unmarshal llm-api response body", slog.String("error", err.Error()))
		return models.AnalyzeData{}, err
	}

	log.Debug("llm-api response", slog.Any("resp", llmResponse))

	content := llmResponse.Choices[0].Message.Content
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var analyzeData models.AnalyzeData
	err = json.Unmarshal([]byte(content), &analyzeData)
	if err != nil {
		log.Error("Failed to unmarshal llm-api response content", slog.String("error", err.Error()))
		return models.AnalyzeData{}, err
	}

	return analyzeData, nil
}
