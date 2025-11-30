package rest

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Util787/mws-content-registry/internal/common"
	"github.com/Util787/mws-content-registry/internal/models"
	"github.com/gin-gonic/gin"
)

type MWSTablesUsecase interface {
	AddYTVideoByURL(url string) error
	AddRecentYTVideos() error
	AddLLMContentAnalyze(recordId string) error
	TakeRecords(pageNum int, pageSize int, sort map[string]string, recordId string, fields []string) ([]models.MWSTableRecord, error)
}

type AiChatUsecase interface {
	SendMessageToChat(chatId int64, message string) (models.Message, error)
	GetChatHistory(chatId int64) ([]models.Message, error)
}

type Handler struct {
	log *slog.Logger
	MWSTablesUsecase
	AiChatUsecase
}

func (h *Handler) addYTVideoByURL(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, log, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := h.MWSTablesUsecase.AddYTVideoByURL(request.URL); err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to add YouTube video", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "YouTube video added successfully"})
}

// Handler for AddRecentYTVideos
func (h *Handler) addRecentYTVideos(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	if err := h.MWSTablesUsecase.AddRecentYTVideos(); err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to add recent YouTube videos", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recent YouTube videos added successfully"})
}

// Handler for AddLLMContentAnalyze
func (h *Handler) addLLMContentAnalyze(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	recordId := c.Param("recordId")
	if recordId == "" {
		newErrorResponse(c, log, http.StatusBadRequest, "record_id is required", errors.New("empty record_id"))
		return
	}

	if err := h.MWSTablesUsecase.AddLLMContentAnalyze(recordId); err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to analyze content with llm", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content analyzed by llm successfully"})
}

func (h *Handler) takeRecords(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	var queryParams struct {
		PageNum  int               `form:"pageNum" binding:"required"`
		PageSize int               `form:"pageSize" binding:"required"`
		Sort     map[string]string `form:"sort"`
		RecordId string            `form:"recordId"`
		Fields   []string          `form:"fields"`
	}

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		newErrorResponse(c, log, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	records, err := h.MWSTablesUsecase.TakeRecords(
		queryParams.PageNum,
		queryParams.PageSize,
		queryParams.Sort,
		queryParams.RecordId,
		queryParams.Fields,
	)
	if err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to retrieve records", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": records})
}

func (h *Handler) sendMessageToChat(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	var request struct {
		ChatId  int64  `json:"chat_id" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, log, http.StatusBadRequest, "Invalid request", err)
		return
	}

	answer, err := h.AiChatUsecase.SendMessageToChat(request.ChatId, request.Message)
	if err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to send message to chat", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer})
}

func (h *Handler) getChatHistory(c *gin.Context) {
	log := common.LogOpAndReqId(c.Request.Context(), common.GetOperationName(), h.log)

	chatIdStr := c.Param("chatId")
	if chatIdStr == "" {
		newErrorResponse(c, log, http.StatusBadRequest, "chatId is required", errors.New("empty chatId"))
		return
	}

	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		newErrorResponse(c, log, http.StatusBadRequest, "invalid chatId", err)
		return
	}

	history, err := h.AiChatUsecase.GetChatHistory(chatId)
	if err != nil {
		newErrorResponse(c, log, http.StatusInternalServerError, "Failed to retrieve chat history", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"chatHistory": history})
}
