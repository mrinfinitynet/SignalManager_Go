package web

import (
	"SignalManager/internal/service/webhooks"

	"github.com/gin-gonic/gin"
)

type BotController struct {
	webhookService *webhooks.WebhookService
}

func NewBotController() *BotController {
	return &BotController{
		webhookService: webhooks.NewWebhookService(),
	}
}

func (ctrl *BotController) Webhook(c *gin.Context) {
	ctrl.webhookService.HandleWebhook(c)
}
