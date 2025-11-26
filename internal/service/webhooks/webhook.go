package webhooks

import (
	"SignalManager/internal/service/webhooks/handler"
	"SignalManager/pkg/logger"
	"encoding/json"
	"io"
	"fmt"
	"math"
	"strings"

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

	// Handle callback queries (button clicks)
	if callbackQuery, ok := data["callback_query"].(map[string]any); ok {
		// CONST
		callbackData, _ := callbackQuery["data"].(string)
		msg, _ := callbackQuery["message"].(map[string]any)
		chat, _ := msg["chat"].(map[string]any)

		chatIdFloat, _ := chat["id"].(float64)
		chatId := fmt.Sprintf("%.0f", chatIdFloat)

		messageIdFloat, _ := msg["message_id"].(float64)
		messageId := int(messageIdFloat)

		dataArray := map[string]interface{}{
			"chat_id":             chatId,
			"message_id":          messageId,
			"callback_data":        callbackData,
		}

		logger.Info("%v", dataArray)

		s.callbackRouter.Handle(c, dataArray)
		return
	}

	// Handle regular messages
	if message, ok := data["message"].(map[string]any); ok {
		// CONST
		chat, _            := message["chat"].(map[string]any)
		chatIdFloat, _     := chat["id"].(float64)
		chatId             := fmt.Sprintf("%.0f", chatIdFloat)
		forwardID          := getForwardID(message)

		// Text 
		caption, _ := message["caption"].(string)
		text, _ := message["text"].(string)
		if text == "" {
			text = caption
		}
		text = strings.TrimSpace(text)

		dataArray := map[string]interface{}{
			"chat_id":         chatId,
			"forward_id":      forwardID,
			"text":            text,
		}

		logger.Info("%v", dataArray)

		s.messageRouter.Handle(c, dataArray)
		return
	}

	// Unknown webhook type
	logger.Warning("[Webhook] Unknown webhook type")
	c.JSON(200, gin.H{"status": "ok"})
}


func getForwardID(message map[string]any) string {
    forwardOrigin, ok := message["forward_origin"].(map[string]any)
    if !ok {
        return ""
    }

    // Try sender_user.id
    if sender, ok := forwardOrigin["sender_user"].(map[string]any); ok {
        if v, ok := sender["id"].(float64); ok {
            return fmt.Sprintf("%.0f", math.Abs(v))
        }
    }

    // Try chat.id
    if chat, ok := forwardOrigin["chat"].(map[string]any); ok {
        if v, ok := chat["id"].(float64); ok {
            return fmt.Sprintf("%.0f", math.Abs(v))
        }
    }

    return ""
}
