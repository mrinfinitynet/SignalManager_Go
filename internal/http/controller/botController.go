package web

import (
	"SignalManager/internal/service/webhooks"
	"net/http"

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
	err := ctrl.webhookService.HandleWebhook(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Webhook received",
	})
}
