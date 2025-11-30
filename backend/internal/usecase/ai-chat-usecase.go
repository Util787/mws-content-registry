package usecase

import (
	"log/slog"
	"time"

	"github.com/Util787/mws-content-registry/internal/common"
	"github.com/Util787/mws-content-registry/internal/models"
)

const defaultPageSize = 20

// returns answer from LLM
func (au *AiChatUsecase) SendMessageToChat(chatId int64, message string) (models.Message, error) {
	userMsgTime := time.Now().Unix()

	sortParams := map[string]string{
		"sort[0][field]": "published_at",
		"sort[0][order]": "desc",
	}

	recentRecords, err := au.MWSTablesClient.TakeRecords(1, defaultPageSize, sortParams, "", []string{})
	if err != nil {
		return models.Message{}, err
	}

	answer, err := au.LLMClient.GenerateChatAnswer(recentRecords, message)
	if err != nil {
		return models.Message{}, err
	}
	answerMsgTime := time.Now().Unix()

	userMessage := models.Message{
		ChatId:    chatId,
		IsUser:    true,
		Message:   message,
		CreatedAt: userMsgTime,
	}
	aiMessage := models.Message{
		ChatId:    chatId,
		IsUser:    false,
		Message:   answer,
		CreatedAt: answerMsgTime,
	}

	au.queueToSave <- userMessage
	au.queueToSave <- aiMessage

	return aiMessage, nil
}

func (au *AiChatUsecase) startMessageSaver() {
	log := au.logger.With("op", common.GetOperationName())
	go func() {
		for mes := range au.queueToSave {
			err := au.MessageStorage.SaveMessage(mes)
			if err != nil {
				log.Warn("failed to save message in storage", slog.String("error", err.Error()))
				continue
			}
		}
	}()
}

func (au *AiChatUsecase) GetChatHistory(chatId int64) ([]models.Message, error) {
	return au.MessageStorage.GetChatHistory(chatId)
}
