package webhooks

import (
	"SignalManager/internal/service/webhooks/handler"
	"SignalManager/pkg/logger"
	"bytes"
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
	var data map[string]any
	var webhookType string

	// Read the request body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("[Webhook] Failed to read body: %v", err)
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	// Log the raw body
	logger.Info("[Webhook] Body: %s", string(bodyBytes))

	// Restore the body for later use
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Parse JSON body into data map
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		logger.Error("[Webhook] Failed to parse JSON: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Check if type is in query parameter first
	if queryType := c.Query("type"); queryType != "" {
		webhookType = queryType
	} else {
		// Get type from JSON body
		typeVal, ok := data["type"].(string)
		if !ok {
			logger.Error("[Webhook] Missing or invalid type field")
			c.JSON(400, gin.H{"error": "Missing type field"})
			return
		}
		webhookType = typeVal
	}

	logger.Info("[Webhook] Type: %s", webhookType)

	switch webhookType {
	case "callback":
		// Handle callback type
		s.callbackRouter.Handle(c, data)
	case "msg":
		// Handle message type
		s.messageRouter.Handle(c, data)
	default:
		// Handle other types
		c.JSON(200, gin.H{
			"message": "else",
		})
	}
}
