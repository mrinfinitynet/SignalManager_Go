package webhooks

import (
	"SignalManager/internal/service/webhooks/handler"
	"SignalManager/pkg/logger"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type WebhookService struct {
	callbackRouter *handler.CallbackRouterService
	messageRouter  *handler.MessageRouterService
}

func NewWebhookService() *WebhookService {
	return &WebhookService{
		callbackRouter: handler.NewCallbackRouterService(),
		messageRouter:  handler.NewMessageRouterService(),
	}
}

func (s *WebhookService) HandleWebhook(c *gin.Context) {
	// Read the request body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("[Webhook] Failed to read body: %v", err)
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	// Parse JSON body into data map
	var data map[string]any
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		logger.Error("[Webhook] Failed to parse JSON: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	logger.Info("Chat Data: %v", data)

	// CONST
	// chat, _ := message["chat"].(map[string]any)
	// text, _ := message["text"].(string)
	// chatId, _ := chat["id"].(float64)

	// Handle callback queries (button clicks)
	if callbackQuery, ok := data["callback_query"].(map[string]any); ok {
		s.callbackRouter.Handle(c, callbackQuery)
		return
	}

	// Handle regular messages
	if message, ok := data["message"].(map[string]any); ok {
		s.messageRouter.Handle(c, message)
		return
	}

	// Unknown webhook type
	logger.Warning("[Webhook] Unknown webhook type")
	c.JSON(200, gin.H{"status": "ok"})
}
