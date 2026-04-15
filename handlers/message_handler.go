package handlers

import (
	"log"
	"net/http"

	"ai-customer-support/models"
	"ai-customer-support/services"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	aiService *services.AIService
}

func NewMessageHandler(aiService *services.AIService) *MessageHandler {
	return &MessageHandler{
		aiService: aiService,
	}
}

func (h *MessageHandler) HandleMessage(c *gin.Context) {
	var req models.MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	log.Printf("Received message from User %s: %s", req.UserID, req.Message)

	aiResp, err := h.aiService.ProcessMessage(c.Request.Context(), req.Message)
	if err != nil {
		log.Printf("AI Service Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process message with AI"})
		return
	}

	log.Printf("AI Response: %s (Category: %s)", aiResp.Response, aiResp.Category)

	c.JSON(http.StatusOK, models.MessageResponse{
		Response: aiResp.Response,
		Category: aiResp.Category,
	})
}
