package webhooks

import (
	"SignalManager/internal/service/webhooks/handler"

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

	// Check if type is in query parameter
	if queryType := c.Query("type"); queryType != "" {
		webhookType = queryType
		// Try to bind JSON body if available
		c.ShouldBindJSON(&data)
	} else {
		// Try to get type from JSON body
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		typeVal, ok := data["type"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"error": "Missing or invalid type field",
			})
			return
		}
		webhookType = typeVal
	}

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
