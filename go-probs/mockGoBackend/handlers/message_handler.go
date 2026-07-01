package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mockGoBackend/services"
)

type MessageHandler struct {
	messages *services.MessageService
}

func NewMessageHandler(messages *services.MessageService) *MessageHandler {
	return &MessageHandler{messages: messages}
}

type createMessageRequest struct {
	Body string `json:"body" binding:"required"`
}

type batchMessagesRequest struct {
	UserIDs []uint `json:"user_ids" binding:"required"`
}

func (h *MessageHandler) Create(c *gin.Context) {
	userID, ok := parseIDParam(c, "user_id")
	if !ok {
		return
	}

	var req createMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.messages.CreateMessage(c.Request.Context(), userID, req.Body)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, msg)
}

// BatchGet demonstrates the concurrent service call — one HTTP request, many DB reads in parallel.
func (h *MessageHandler) BatchGet(c *gin.Context) {
	var req batchMessagesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := h.messages.GetMessagesForUsers(c.Request.Context(), req.UserIDs)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, results)
}
